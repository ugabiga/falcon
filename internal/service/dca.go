package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/common/timer"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
	"math"
)

var (
	ErrExchangeNotFound = errors.New("exchange_not_found")
)

type TaskOrderInfo struct {
	TaskID           string
	TradingAccountID string
	UserID           string
}

type DcaService struct {
	repo       *repository.DynamoRepository
	encryption *encryption.Encryption
}

func NewDcaService(
	repo *repository.DynamoRepository,
	encryption *encryption.Encryption,
) *DcaService {
	return &DcaService{
		repo:       repo,
		encryption: encryption,
	}
}

func (s DcaService) GetTarget() ([]TaskOrderInfo, error) {
	ctx := context.Background()
	now := timer.NoSeconds()

	log.Printf("Searching for tasks with next execution time: %s", now.String())

	// TODO add pagination
	tasks, err := s.repo.GetTasksByActiveNextExecutionTime(ctx, now)
	if err != nil {
		return nil, err
	}
	log.Printf("Tasks: %+v", debug.ToJSONInlineStr(tasks))

	var taskOrderInfos []TaskOrderInfo
	for _, t := range tasks {
		taskOrderInfos = append(taskOrderInfos, TaskOrderInfo{
			TaskID:           t.ID,
			TradingAccountID: t.TradingAccountID,
			UserID:           t.UserID,
		})
	}

	return taskOrderInfos, nil
}

func (s DcaService) Order(orderInfo TaskOrderInfo) error {
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

func (s DcaService) createTaskHistory(ctx context.Context, orderErr error, t *model.Task) error {
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
func (s DcaService) updateNextTaskExecutionTime(ctx context.Context, userID string, t *model.Task) error {
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

func (s DcaService) OrderFromBinance(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	symbol := t.Symbol + t.Currency
	size := t.Size
	key := tradingAccount.Key
	decryptedSecret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}
	log.Printf("order at binance: key: %s, size: %f, symbol: %s", key, size, symbol)

	c := client.NewBinanceClient(key, decryptedSecret, false)

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		log.Printf("Error getting ticker: %s", err.Error())
		return err
	}
	if ticker == nil {
		return ErrTickerNotFound
	}

	tickerPriceDecimalCount := str.New(ticker.Price).CountDecimalCount()
	tickerPrice := str.New(ticker.Price).ToFloat64Default(0)
	roundedTickerPrice := math.Round(tickerPrice*math.Pow10(tickerPriceDecimalCount)) / math.Pow10(tickerPriceDecimalCount)

	order, err := c.PlaceOrderAtPrice(ctx,
		symbol,
		client.HoldSideLong,
		str.FromFloat64(size).Val(),
		str.FromFloat64(roundedTickerPrice).Val(),
	)
	if err != nil {
		return err
	}

	log.Printf("order: %+v", debug.ToJSONInlineStr(order))

	return nil
}

func (s DcaService) OrderFromUpbit(
	ctx context.Context,
	tradingAccount *model.TradingAccount,
	t *model.Task,
) error {
	symbol := t.Currency + "-" + t.Symbol
	size := t.Size
	key := tradingAccount.Key
	decryptedSecret, err := s.encryption.Decrypt(tradingAccount.Secret)
	if err != nil {
		return err
	}

	log.Printf("order at upbit: key: %s, size: %f, symbol: %s", key, size, symbol)

	c := client.NewUpbitClient(key, decryptedSecret)

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		log.Printf("Error getting ticker: %s", err.Error())
		return err
	}

	if ticker == nil {
		return ErrTickerNotFound
	}
	log.Printf("ticker: %+v", debug.ToJSONInlineStr(ticker))

	tradePrice := ticker.TradePrice
	tradePriceStr := str.FromFloat64(tradePrice).Val()
	sizeStr := str.FromFloat64(size).Val()

	order, err := c.PlaceLongOrderAt(ctx, symbol, sizeStr, tradePriceStr)
	if err != nil {
		return err
	}

	log.Printf("order: %+v", debug.ToJSONInlineStr(order))

	return nil
}
