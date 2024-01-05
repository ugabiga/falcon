package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/common/timer"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
)

var (
	ErrTickerNotFound = errors.New("ticker not found")
)

type TaskOrderInfo struct {
	TaskID           string
	TradingAccountID string
	UserID           string
}

type DcaService struct {
	db          *ent.Client
	userRepo    *repository.UserDynamoRepository
	tradingRepo *repository.TradingDynamoRepository
	encryption  *encryption.Encryption
}

func NewDcaService(
	db *ent.Client,
	userRepo *repository.UserDynamoRepository,
	tradingRepo *repository.TradingDynamoRepository,
	encryption *encryption.Encryption,
) *DcaService {
	return &DcaService{
		db:          db,
		userRepo:    userRepo,
		tradingRepo: tradingRepo,
		encryption:  encryption,
	}
}

func (s DcaService) GetTarget() ([]TaskOrderInfo, error) {
	ctx := context.Background()
	now := timer.NoSeconds()

	log.Printf("Searching for tasks with next execution time: %s", now.String())

	// TODO add pagination
	tasks, err := s.tradingRepo.GetTasksByActiveNextExecutionTime(ctx, now)
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d tasks", len(tasks))

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

	tradingAccount, err := s.tradingRepo.GetTradingAccount(ctx, orderInfo.UserID, orderInfo.TradingAccountID)
	if err != nil {
		return err
	}

	t, err := s.tradingRepo.GetTask(ctx, orderInfo.TradingAccountID, orderInfo.TaskID)
	if err != nil {
		return err
	}

	switch tradingAccount.Exchange {
	case "upbit":
		err = s.orderUpbitAt(
			ctx,
			tradingAccount,
			t,
		)
	default:
		return errors.New("exchange not found")
	}

	if err != nil {
		return err
	}

	return s.updateNextTaskExecutionTime(ctx, orderInfo.UserID, t)
}
func (s DcaService) updateNextTaskExecutionTime(ctx context.Context, userID string, t *model.Task) error {
	u, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return err
	}

	nextCronExecutionTime, err := nextCronExecutionTime(t.Cron, u.Timezone)
	if err != nil {
		return err
	}
	t.NextExecutionTime = nextCronExecutionTime

	_, err = s.tradingRepo.UpdateTask(ctx, *t)
	if err != nil {
		return err
	}

	return nil
}

func (s DcaService) orderUpbitAt(
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

	log.Printf("key: %s, size: %f, symbol: %s", key, size, symbol)
	return nil

	c := client.NewUpbitClient(key, decryptedSecret)

	ticker, err := c.Ticker(ctx, symbol)
	if err != nil {
		return err
	}

	if ticker == nil {
		return ErrTickerNotFound
	}

	tradePrice := ticker.TradePrice
	tradePriceStr := str.FromFloat64(tradePrice).Val()
	sizeStr := str.FromFloat64(size).Val()

	order, err := c.PlaceLongOrderAt(ctx, symbol, sizeStr, tradePriceStr)
	if err != nil {
		return err
	}

	log.Printf("order: %+v", order)

	return nil
}
