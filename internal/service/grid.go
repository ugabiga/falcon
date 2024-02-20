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
	taskType := model.TaskTypeLongGrid
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
		orderErr := s.OrderFromBinanceFuture(
			ctx,
			tradingAccount,
			t,
		)
		if err := s.createTaskHistory(ctx, orderErr, t); err != nil {
			return err
		}
	case model.ExchangeBinanceSpot:
		orderErr := s.OrderFromBinanceSpot(
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

func (s GridService) OrderFromBinanceSpot(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	params, err := t.GridParams()
	if err != nil {
		return err
	}

	symbol := t.Symbol + t.Currency
	size := t.Size
	key := tradingAccount.Key
	decryptedSecret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}
	c := client.NewBinanceSpotClient(key, decryptedSecret, false)
	log.Printf("OrderFromBinanceSpot: key: %s, size: %f, symbol: %s grid params: %+v",
		key,
		size,
		symbol,
		debug.ToJSONInlineStr(params),
	)

	//cancel all open orders
	orders, err := c.OpenPositionOrders(ctx, symbol)
	if err != nil {
		return err
	}

	var orderIDs []int64
	for _, order := range orders {
		orderIDs = append(orderIDs, order.OrderID)
	}

	if len(orderIDs) > 0 {
		_, err = c.CancelOpenOrders(ctx, symbol)
		if err != nil {
			return err
		}
	}

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		log.Printf("Error getting ticker: %s", err.Error())
		return err
	}
	if ticker == nil {
		return ErrTickerNotFound
	}

	tickSizeStr, _, err := c.TickAndStepSize(ctx, symbol)
	if err != nil {
		log.Printf("Error getting lotSize: %s", err.Error())
		return err
	}
	tickSize := str.New(tickSizeStr).ToFloat64Default(0)
	tickerPriceDecimalCount := str.New(ticker.Price).CountDecimalCount()
	tickerPrice := str.New(ticker.Price).ToFloat64Default(0)

	for i := int64(0); i < params.Quantity; i++ {
		percentDownTickerPrice := tickerPrice - (tickerPrice * params.GapPercent / 100)
		roundedTickerPrice := math.Round(percentDownTickerPrice*math.Pow10(tickerPriceDecimalCount)) / math.Pow10(tickerPriceDecimalCount)
		trimmedPrice := math.Round(roundedTickerPrice/tickSize) * tickSize
		log.Printf("tickerPrice: %f,  roundedTickerPrice: %f, trimmedPrice: %f", tickerPrice, roundedTickerPrice, trimmedPrice)

		order, err := c.PlaceOrderAtPrice(ctx,
			symbol,
			client.HoldSideLong,
			str.FromFloat64(size).Val(),
			str.FromFloat64(trimmedPrice).Val(),
		)
		if err != nil {
			log.Printf("Error placing order: %s", err)
			continue
		}
		log.Printf("Successfully placed order: %+v", debug.ToJSONInlineStr(order))

		tickerPrice = percentDownTickerPrice
	}

	return nil
}

func (s GridService) OrderFromBinanceFuture(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	params, err := t.GridParams()
	if err != nil {
		return err
	}

	symbol := t.Symbol + t.Currency
	size := t.Size
	key := tradingAccount.Key
	decryptedSecret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}
	c := client.NewBinanceFutureClient(key, decryptedSecret, false)
	log.Printf("OrderFromBinanceFuture: key: %s, size: %f, symbol: %s grid params: %+v",
		key,
		size,
		symbol,
		debug.ToJSONInlineStr(params),
	)

	//cancel all open orders
	orders, err := c.OpenPositionOrders(ctx, symbol)
	if err != nil {
		return err
	}

	var orderIDs []int64
	for _, order := range orders {
		orderIDs = append(orderIDs, order.OrderID)
	}

	if len(orderIDs) > 0 {
		_, err = c.CancelOpenOrders(ctx, symbol, orderIDs)
		if err != nil {
			return err
		}
	}

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		log.Printf("Error getting ticker: %s", err.Error())
		return err
	}
	if ticker == nil {
		return ErrTickerNotFound
	}

	tickSizeStr, _, err := c.TickAndStepSize(ctx, symbol)
	if err != nil {
		log.Printf("Error getting lotSize: %s", err.Error())
		return err
	}
	tickSize := str.New(tickSizeStr).ToFloat64Default(0)
	tickerPriceDecimalCount := str.New(ticker.Price).CountDecimalCount()
	tickerPrice := str.New(ticker.Price).ToFloat64Default(0)

	for i := int64(0); i < params.Quantity; i++ {
		percentDownTickerPrice := tickerPrice - (tickerPrice * params.GapPercent / 100)
		roundedTickerPrice := math.Round(percentDownTickerPrice*math.Pow10(tickerPriceDecimalCount)) / math.Pow10(tickerPriceDecimalCount)
		trimmedPrice := math.Round(roundedTickerPrice/tickSize) * tickSize
		log.Printf("tickerPrice: %f,  roundedTickerPrice: %f, trimmedPrice: %f", tickerPrice, roundedTickerPrice, trimmedPrice)

		order, err := c.PlaceOrderAtPrice(ctx,
			symbol,
			client.HoldSideLong,
			str.FromFloat64(size).Val(),
			str.FromFloat64(trimmedPrice).Val(),
		)
		if err != nil {
			log.Printf("Error placing order: %s", err)
			continue
		}
		log.Printf("Successfully placed order: %+v", debug.ToJSONInlineStr(order))

		tickerPrice = percentDownTickerPrice
	}

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

	orders, err := c.Orders(ctx, symbol)
	if err != nil {
		return err
	}

	for _, order := range orders {
		if _, err := c.CancelOrder(ctx, order.UUID); err != nil {
			log.Printf("Error canceling order: %s", err.Error())
		}
	}

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
