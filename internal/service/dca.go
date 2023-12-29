package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/str"
	"github.com/ugabiga/falcon/internal/ent"
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
	Symbol     string
	Size       float64
	Exchange   string
	Identifier string
	Credential string
}

func (s DcaService) GetTarget() ([]TaskOrderInfo, error) {
	ctx := context.Background()
	tasks, err := s.db.Task.Query().WithTradingAccount().All(ctx)
	if err != nil {
		return nil, err
	}

	var taskOrderInfos []TaskOrderInfo
	for _, task := range tasks {
		if task.Edges.TradingAccount != nil {
			taskOrderInfos = append(taskOrderInfos, TaskOrderInfo{
				Symbol:     task.CryptoCurrency,
				Size:       task.Size,
				Exchange:   task.Edges.TradingAccount.Exchange,
				Identifier: task.Edges.TradingAccount.Identifier,
				Credential: task.Edges.TradingAccount.Credential,
			})
		}
	}

	return taskOrderInfos, nil
}

func (s DcaService) Order(orderInfo TaskOrderInfo) error {
	switch orderInfo.Exchange {
	case "upbit":
		return s.orderUpbitAt(
			context.Background(),
			orderInfo.Symbol,
			orderInfo.Size,
			orderInfo.Identifier,
			orderInfo.Credential,
		)
	default:
		return errors.New("exchange not found")
	}
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
