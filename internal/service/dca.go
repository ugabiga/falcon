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

func (s DcaService) orderUpbitAt(
	ctx context.Context,
	symbol string,
	size float64,
	key string,
	secret string,
) error {
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
