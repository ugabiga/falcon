package client

import (
	"context"
	"errors"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/str"
	"log"
)

type BinanceSpotClient struct {
	spot *binance.Client
}

func NewBinanceSpotClient(apiKey, secretKey string, isTest bool) *BinanceSpotClient {
	return &BinanceSpotClient{
		spot: binance.NewClient(apiKey, secretKey),
	}
}

func (c *BinanceSpotClient) MinQuantity(ctx context.Context, symbol string) (string, error) {
	resp, err := c.spot.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return "", err
	}

	for _, s := range resp.Symbols {
		if s.Symbol == symbol {
			lotSizeFilter := s.PriceFilter()
			log.Printf("Symbol: %+v", debug.ToJSONStr(s))
			return lotSizeFilter.MinPrice, nil
		}
	}

	return "", err
}

func (c *BinanceSpotClient) TickAndStepSize(ctx context.Context, symbol string) (string, string, error) {
	resp, err := c.spot.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return "", "", err
	}

	for _, s := range resp.Symbols {
		if s.Symbol == symbol {
			lotSizeFilter := s.LotSizeFilter()
			priceFilter := s.PriceFilter()
			log.Printf("Symbol: %+v", debug.ToJSONStr(s))
			return priceFilter.TickSize, lotSizeFilter.StepSize, nil
		}
	}

	return "", "", err
}

func (c *BinanceSpotClient) LotSize(ctx context.Context, symbol string) (*binance.LotSizeFilter, error) {
	resp, err := c.spot.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, s := range resp.Symbols {
		if s.Symbol == symbol {
			log.Printf("Symbol: %+v", debug.ToJSONStr(s))
			lotSizeFilter := s.LotSizeFilter()
			return lotSizeFilter, nil
		}
	}

	return nil, nil

}

//func (c *BinanceSpotClient) Balance(ctx context.Context) (*binance.Balance, error) {
//	resp, err := c.spot.NewGetBalanceService().Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, b := range resp {
//		if b.Asset == "USDT" {
//			return b, nil
//		}
//	}
//
//	return nil, nil
//
//}

func (c *BinanceSpotClient) Ticker(ctx context.Context, symbol string) (*binance.SymbolPrice, error) {
	resp, err := c.spot.NewListPricesService().Symbol(symbol).Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range resp {
		if t.Symbol == symbol {
			return t, nil
		}
	}

	return nil, nil
}

//func (c *BinanceSpotClient) PositionWithoutSideIncludeEmpty(ctx context.Context, symbol string) (*futures.AccountPosition, error) {
//	resp, err := c.spot.NewGetAccountService().Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, p := range resp.Balances {
//		if p.Symbol == symbol {
//			return p, nil
//		}
//	}
//
//	return nil, nil
//}

//func (c *BinanceSpotClient) PositionWithoutSide(ctx context.Context, symbol string) (*futures.AccountPosition, error) {
//	resp, err := c.spot.NewGetAccountService().Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, p := range resp.Positions {
//		if p.Symbol == symbol {
//			positionAmt := str.New(p.PositionAmt).ToFloat64Default(0)
//			if positionAmt == 0 {
//				continue
//			}
//
//			return p, nil
//		}
//	}
//
//	return nil, nil
//}

func (c *BinanceSpotClient) PositionSide(p futures.AccountPosition) (string, error) {
	positionAmt := str.New(p.PositionAmt).ToFloat64Default(0)
	if positionAmt == 0 {
		return "", errors.New("position amount is 0")
	}

	if positionAmt > 0 {
		return HoldSideLong, nil
	} else if positionAmt < 0 {
		return HoldSideShort, nil
	}

	return "", ErrPositionNotFound
}

//func (c *BinanceSpotClient) PositionWithEmptyValue(ctx context.Context, symbol, holdSide string) (*futures.AccountPosition, error) {
//	resp, err := c.spot.NewGetAccountService().Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, p := range resp.Positions {
//		if p.Symbol == symbol {
//			return p, nil
//		}
//	}
//
//	return nil, nil
//}

//func (c *BinanceSpotClient) Position(ctx context.Context, symbol, holdSide string) (*futures.AccountPosition, error) {
//	resp, err := c.spot.NewGetAccountService().Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, p := range resp.Positions {
//		if p.Symbol == symbol {
//			positionAmt := str.New(p.PositionAmt).ToFloat64Default(0)
//			if positionAmt == 0 {
//				continue
//			}
//
//			side := HoldSideShort
//			if positionAmt > 0 {
//				side = HoldSideLong
//			}
//
//			if side == holdSide {
//				return p, nil
//			}
//		}
//	}
//
//	return nil, nil
//}

func (c *BinanceSpotClient) PlaceOrderAtPrice(ctx context.Context, symbol, holdSide, size, price string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeLimit).
		Side(c.convertHoldSide(holdSide)).
		Quantity(size).
		Price(price).
		TimeInForce(binance.TimeInForceTypeGTC).
		Do(ctx)
	if err != nil {
		return nil, c.errorConverter(err)
	}

	return resp, nil
}

func (c *BinanceSpotClient) PlaceOrder(ctx context.Context, symbol, holdSide, size string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeMarket).
		Side(c.convertHoldSide(holdSide)).
		Quantity(size).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetTP TODO : Need to check it works
func (c *BinanceSpotClient) SetTP(ctx context.Context, symbol, holdSide, price string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeTakeProfit).
		Side(c.oppositeHoldSide(holdSide)).
		StopPrice(price).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetTPLimit TODO : Need to check it works
func (c *BinanceSpotClient) SetTPLimit(ctx context.Context, symbol, holdSide, price, size string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeTakeProfitLimit).
		Side(c.oppositeHoldSide(holdSide)).
		Price(price).
		StopPrice(price).
		Quantity(size).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetSL TODO : Need to check it works
func (c *BinanceSpotClient) SetSL(ctx context.Context, symbol, holdSide, price string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeStopLoss).
		Side(c.oppositeHoldSide(holdSide)).
		StopPrice(price).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetSLLimit TODO : Need to check it works
func (c *BinanceSpotClient) SetSLLimit(ctx context.Context, symbol, holdSide, price, size string) (*binance.CreateOrderResponse, error) {
	resp, err := c.spot.NewCreateOrderService().
		Symbol(symbol).
		Type(binance.OrderTypeStopLossLimit).
		Side(c.oppositeHoldSide(holdSide)).
		Price(price).
		StopPrice(price).
		Quantity(size).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceSpotClient) OpenTPSLTSOrders(ctx context.Context, symbol string) ([]*binance.Order, error) {
	resp, err := c.spot.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]*binance.Order, 0)
	for _, order := range resp {
		if order.Type == binance.OrderTypeStopLoss ||
			order.Type == binance.OrderTypeTakeProfit {
			orders = append(orders, order)
		}
	}

	return orders, nil
}

func (c *BinanceSpotClient) OpenPositionOrders(ctx context.Context, symbol string) ([]*binance.Order, error) {
	resp, err := c.spot.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]*binance.Order, 0)
	for _, order := range resp {
		if order.Type == binance.OrderTypeLimit ||
			order.Type == binance.OrderTypeMarket {
			orders = append(orders, order)
		}
	}

	return orders, nil
}

// CancelOpenOrders It cancels all open orders
func (c *BinanceSpotClient) CancelOpenOrders(ctx context.Context, symbol string) (*binance.CancelOpenOrdersResponse, error) {
	cancelResult, err := c.spot.NewCancelOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return cancelResult, nil
}

func (c *BinanceSpotClient) LimitOrder(ctx context.Context, symbol string) (*binance.Order, error) {
	resp, err := c.spot.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, order := range resp {
		if order.Type == binance.OrderTypeLimit &&
			order.Symbol == symbol {
			return order, nil
		}
	}

	return nil, nil
}

func (c *BinanceSpotClient) oppositeHoldSide(holdSide string) binance.SideType {
	side := c.convertHoldSide(holdSide)
	switch side {
	case binance.SideTypeBuy:
		return binance.SideTypeSell
	case binance.SideTypeSell:
		return binance.SideTypeBuy
	default:
		return side
	}
}

func (c *BinanceSpotClient) convertHoldSide(holdSide string) binance.SideType {
	switch holdSide {
	case HoldSideLong:
		return binance.SideTypeBuy
	case HoldSideShort:
		return binance.SideTypeSell
	default:
		return binance.SideType(holdSide)
	}
}

func (c *BinanceSpotClient) errorConverter(apiErr error) error {
	var e *common.APIError
	switch {
	case errors.As(apiErr, &e):
		switch e.Code {
		case -4164:
			return ErrMinNotional
		case -1111:
			return ErrPrecisionError
		default:
			return apiErr
		}
	default:
		return apiErr
	}
}
