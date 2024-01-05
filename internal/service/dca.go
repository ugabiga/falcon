package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/common/timer"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/task"
	"github.com/ugabiga/falcon/internal/ent/user"
	"log"
)

var (
	ErrTickerNotFound = errors.New("ticker not found")
)

type TaskOrderInfo struct {
	TaskID   int
	Cron     string
	Symbol   string
	Currency string
	Size     float64
	Exchange string
	Key      string
	Secret   string
}

type DcaService struct {
	db         *ent.Client
	encryption *encryption.Encryption
}

func NewDcaService(
	db *ent.Client,
	encryption *encryption.Encryption,
) *DcaService {
	return &DcaService{
		db:         db,
		encryption: encryption,
	}
}

func (s DcaService) GetTarget() ([]TaskOrderInfo, error) {
	ctx := context.Background()
	now := timer.NowNoMinuteAndSeconds()

	log.Printf("Searching for tasks with next execution time: %s", now.String())

	// TODO add pagination
	tasks, err := s.db.Task.Query().
		Where(
			task.NextExecutionTime(now),
			task.IsActive(true),
		).
		WithTradingAccount().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var taskOrderInfos []TaskOrderInfo
	for _, t := range tasks {
		if t.Edges.TradingAccount != nil {
			taskOrderInfos = append(taskOrderInfos, TaskOrderInfo{
				TaskID:   t.ID,
				Cron:     t.Cron,
				Symbol:   t.Symbol,
				Currency: t.Currency,
				Size:     t.Size,
				Exchange: t.Edges.TradingAccount.Exchange,
				Key:      t.Edges.TradingAccount.Key,
				Secret:   t.Edges.TradingAccount.Secret,
			})
		}
	}

	return taskOrderInfos, nil
}

func (s DcaService) Order(orderInfo TaskOrderInfo) error {
	var err error

	switch orderInfo.Exchange {
	case "upbit":
		err = s.orderUpbitAt(
			context.Background(),
			orderInfo,
		)
	default:
		return errors.New("exchange not found")
	}

	if err != nil {
		return err
	}

	return s.updateNextTaskExecutionTime(orderInfo)
}
func (s DcaService) updateNextTaskExecutionTime(orderInfo TaskOrderInfo) error {
	ctx := context.Background()
	t, err := s.db.Task.Query().
		Where(
			task.ID(orderInfo.TaskID),
		).
		WithTradingAccount().
		First(ctx)
	if err != nil {
		return err
	}

	u, err := s.db.User.Query().
		Where(
			user.IDEQ(t.Edges.TradingAccount.UserID),
		).
		First(ctx)
	if err != nil {
		return err
	}

	nextCronExecutionTime, err := nextCronExecutionTime(t.Cron, u.Timezone)
	if err != nil {
		return err
	}

	_, err = s.db.Task.UpdateOneID(orderInfo.TaskID).
		SetNextExecutionTime(nextCronExecutionTime).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s DcaService) orderUpbitAt(
	ctx context.Context,
	orderInfo TaskOrderInfo,
) error {
	symbol := orderInfo.Currency + "-" + orderInfo.Symbol
	size := orderInfo.Size
	key := orderInfo.Key
	decryptedSecret, err := s.encryption.Decrypt(orderInfo.Secret)
	if err != nil {
		return err
	}

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
