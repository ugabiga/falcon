package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/client/upbit"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/common/timer"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
	"math"
	"strconv"
)

type GridService struct {
	repo       *repository.DynamoRepository
	encryption *encryption.Encryption
}

func NewGridService(
	repo *repository.DynamoRepository,
	encryption *encryption.Encryption,
) *GridService {
	return &GridService{
		repo:       repo,
		encryption: encryption,
	}
}

func (s GridService) GetTarget(t Timer) ([]TaskOrderInfo, error) {
	ctx := context.Background()
	now := timer.NoSeconds()
	if t != nil {
		now = t.NoSeconds()
	}
	taskType := model.TaskTypeGrid
	tasks, err := s.repo.GetTasksByActiveNextExecutionTimeAndType(ctx, now, taskType)
	if err != nil {
		return nil, err
	}

	var taskOrderInfos []TaskOrderInfo
	for _, t := range tasks {
		taskOrderInfos = append(taskOrderInfos, TaskOrderInfo{
			TaskType:         t.Type,
			TaskID:           t.ID,
			TradingAccountID: t.TradingAccountID,
			UserID:           t.UserID,
		})
	}

	return taskOrderInfos, nil
}

func (s GridService) Order(orderInfo TaskOrderInfo) error {
	ctx := context.Background()
	var err error

	tradingAccount, err := s.repo.GetTradingAccount(ctx, orderInfo.UserID, orderInfo.TradingAccountID)
	if err != nil {
		return err
	}

	t, err := s.repo.GetTask(ctx, orderInfo.TradingAccountID, orderInfo.TaskID)
	if err != nil {
		return err
	}

	switch tradingAccount.Exchange {
	case model.ExchangeUpbit:
		orderErr := s.OrderFromUpbit(
			ctx,
			tradingAccount,
			t,
		)
		if err := s.createTaskHistory(ctx, orderErr, t); err != nil {
			return err
		}
	case model.ExchangeBinanceFutures:
		orderErr := s.OrderFromBinance(
			ctx,
			tradingAccount,
			t,
		)
		if err := s.createTaskHistory(ctx, orderErr, t); err != nil {
			return err
		}
	default:
		orderErr := ErrExchangeNotFound
		if err := s.createTaskHistory(ctx, orderErr, t); err != nil {
			return err
		}
		return orderErr
	}

	if err := s.updateNextTaskExecutionTime(ctx, orderInfo.UserID, t); err != nil {
		return err
	}

	return nil
}

func (s GridService) OrderFromBinance(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	return nil

}

func (s GridService) OrderFromUpbit(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	params, err := t.GridParams()
	if err != nil {
		return err
	}
	log.Printf("Grid params: %+v", debug.ToJSONInlineStr(params))

	symbol := t.Currency + "-" + t.Symbol
	size := t.Size
	key := tradingAccount.Key
	decryptedSecret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}
	log.Printf("OrderFromUpbit: key: %s, size: %f, symbol: %s", key, size, symbol)

	c := client.NewUpbitClient(key, decryptedSecret)

	orderBook, err := c.OrderBook(ctx, symbol)
	if err != nil {
		return err
	}

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		log.Printf("Error getting ticker: %s", err.Error())
		return err
	}
	if ticker == nil {
		return ErrTickerNotFound
	}

	tradePrice := ticker.TradePrice
	tradeUnitPrice := s.upbitUnitPrice(orderBook)
	tradeLot := s.upbitLot(tradeUnitPrice)

	for i := int64(0); i < params.Quantity; i++ {
		percentDownTradePrice := tradePrice - (tradePrice * float64(params.GapPercent) / 100)
		percentDownTradePrice = math.Round(percentDownTradePrice/tradeLot) * tradeLot
		percentDownTradePrice = percentDownTradePrice - math.Mod(percentDownTradePrice, tradeUnitPrice)
		log.Printf("tradePrice: %f,  percentDownTradePrice: %f", tradePrice, percentDownTradePrice)

		tradePriceStr := str.FromFloat64(percentDownTradePrice).Val()
		sizeStr := str.FromFloat64(size).Val()
		order, err := c.PlaceLongOrderAt(ctx, symbol, sizeStr, tradePriceStr)
		if err != nil {
			continue
		}
		log.Printf("order: %+v", debug.ToJSONInlineStr(order))

		tradePrice = percentDownTradePrice
	}

	return nil
}
func (s GridService) upbitLot(unitPrice float64) float64 {
	lot := countZeros(int(unitPrice))

	return math.Pow(10, float64(lot))
}
func (s GridService) upbitUnitPrice(orderBook *upbit.OrderBook) float64 {
	minimumSubtract := int64(0)
	for i, orderBookUnit := range orderBook.OrderbookUnits {
		if i == 0 {
			continue
		}
		previousOrderBookUnit := orderBook.OrderbookUnits[i-1]
		currentOrderBookUnit := orderBookUnit
		subtract := currentOrderBookUnit.AskPrice - previousOrderBookUnit.AskPrice

		if minimumSubtract == 0 {
			minimumSubtract = subtract
		}
		if subtract < minimumSubtract {
			minimumSubtract = subtract
		}
	}

	return float64(minimumSubtract)

}

func countZeros(n int) int {
	str := strconv.Itoa(n)
	count := 0
	for _, char := range str {
		if char == '0' {
			count++
		}
	}
	return count
}

func (s GridService) createTaskHistory(ctx context.Context, orderErr error, t *model.Task) error {
	isSuccess := true
	logMessage := "task executed successfully"
	if orderErr != nil {
		logMessage = orderErr.Error()
		isSuccess = false
	}
	th := model.TaskHistory{
		TaskID:           t.ID,
		TradingAccountID: t.TradingAccountID,
		UserID:           t.UserID,
		IsSuccess:        isSuccess,
		Log:              logMessage,
	}

	log.Printf("Creating task history: %+v", debug.ToJSONInlineStr(th))

	_, err := s.repo.CreateTaskHistory(ctx, th)
	if err != nil {
		return err
	}

	return nil
}

func (s GridService) updateNextTaskExecutionTime(ctx context.Context, userID string, t *model.Task) error {
	u, err := s.repo.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	nextCronExecutionTime, err := nextCronExecutionTime(t.Cron, u.Timezone)
	if err != nil {
		return err
	}
	t.NextExecutionTime = nextCronExecutionTime

	log.Printf("Updating task: %+v", debug.ToJSONInlineStr(t))

	_, err = s.repo.UpdateTask(ctx, *t)
	if err != nil {
		return err
	}

	return nil
}
