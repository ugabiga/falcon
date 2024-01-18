package service

import (
	"context"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/inti"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/common/timer"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
	"math"
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
	tradeUnit := orderBook.UnitPrice()
	tradeLot := s.upbitLot(tradeUnit)

	for i := int64(0); i < params.Quantity; i++ {
		percentDownTradePrice := tradePrice - (tradePrice * float64(params.GapPercent) / 100)
		appliedLotPrice := math.Round(percentDownTradePrice/tradeLot) * tradeLot
		appliedUnitPrice := appliedLotPrice - math.Mod(appliedLotPrice, tradeUnit)
		log.Printf("tradePrice: %f,  percentDownTradePrice: %f", tradePrice, appliedUnitPrice)

		order, err := c.PlaceLongOrderAt(
			ctx,
			symbol,
			str.FromFloat64(size).Val(),
			str.FromFloat64(appliedUnitPrice).Val(),
		)
		if err != nil {
			log.Printf("Error placing order: %s", err.Error())
			continue
		}
		if order.UUID == "" {
			log.Printf("Error placing order: %s", debug.ToJSONInlineStr(order))
			continue
		}
		log.Printf("Successfully placed order: %+v", debug.ToJSONInlineStr(order))

		tradePrice = percentDownTradePrice
	}

	return nil
}
func (s GridService) upbitLot(unitPrice float64) float64 {
	lot := inti.CountZeros(int(unitPrice))
	return math.Pow(10, float64(lot))
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
