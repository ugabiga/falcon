package client

import (
	"context"
	"github.com/ugabiga/falcon/internal/client/upbit"
)

type UpbitClient struct {
	client *upbit.Client
}

func NewUpbitClient(key, secret string) *UpbitClient {
	return &UpbitClient{
		client: upbit.NewUpbitClient(key, secret),
	}
}

func (c *UpbitClient) Accounts() ([]upbit.Account, error) {
	accounts, err := c.client.Accounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *UpbitClient) PlaceLongOrderAt(ctx context.Context, symbol, size, priceInKRW string) (*upbit.CreateOrderResponse, error) {
	order, err := c.client.PlaceLongOrderAt(ctx, symbol, size, priceInKRW)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *UpbitClient) Ticker(ctx context.Context, symbol string) (*upbit.Ticker, error) {
	ticker, err := c.client.Ticker(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (c *UpbitClient) TickerPublic(ctx context.Context, symbol string) (*upbit.Ticker, error) {
	ticker, err := c.client.TickerPublic(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (c *UpbitClient) OrderChance(ctx context.Context, symbol string) (*upbit.OrderChange, error) {
	r, err := c.client.OrderChance(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (c *UpbitClient) OrderBook(ctx context.Context, symbol string) (*upbit.OrderBook, error) {
	r, err := c.client.OrderBook(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *UpbitClient) Orders(ctx context.Context, symbol string) ([]upbit.Order, error) {
	r, err := c.client.Orders(ctx, symbol)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *UpbitClient) CancelOrder(ctx context.Context, uuid string) (*upbit.Order, error) {
	r, err := c.client.CancelOrder(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return r, nil
}
