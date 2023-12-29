package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
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

type DcaService struct {
	db *ent.Client
}

func NewDcaService(
	db *ent.Client,
) *DcaService {
	return &DcaService{
		db: db,
	}
}

type TaskOrderInfo struct {
	TaskID   int
	Cron     string
	Symbol   string
	Size     float64
	Exchange string
	Key      string
	Secret   string
}

func (s DcaService) GetTarget() ([]TaskOrderInfo, error) {
	ctx := context.Background()
	now := timer.NoSeconds()

	log.Printf("Searching for tasks with next execution time: %s", now.String())

	// TODO add pagination
	tasks, err := s.db.Task.Query().
		Where(
			task.NextExecutionTime(now),
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
			orderInfo.Symbol,
			orderInfo.Size,
			orderInfo.Key,
			orderInfo.Secret,
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

	nextCronExecutionTime, err := NextCronExecutionTime(t.Cron, u.Timezone)
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
	symbol string,
	size float64,
	key string,
	secret string,
) error {
	log.Printf("did it")
	return nil

	c := client.NewUpbitClient("", "")

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
