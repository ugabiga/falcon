package client

import (
	"context"
	"errors"
	"github.com/adshao/go-binance/v2/common"
	"github.com/ugabiga/falcon/internal/common/debug"
	"github.com/ugabiga/falcon/internal/common/str"
	"log"

	"github.com/adshao/go-binance/v2/futures"
)

var (
	HoldSideLong  = "long"
	HoldSideShort = "short"
)

var (
	ErrPositionNotFound = errors.New("position not found")
	ErrPrecisionError   = errors.New("precision_error")
	ErrMinNotional      = errors.New("not_satisfied_min_notional")
)

type BinanceClient struct {
	future *futures.Client
}

func NewBinanceClient(apiKey, secretKey string, isTest bool) *BinanceClient {
	if isTest {
		futures.UseTestnet = true
	}
	c := futures.NewClient(apiKey, secretKey)
	return &BinanceClient{
		future: c,
	}
}

func (c *BinanceClient) MinQuantity(ctx context.Context, symbol string) (string, error) {
	resp, err := c.future.NewExchangeInfoService().Do(ctx)
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

func (c *BinanceClient) TickAndStepSize(ctx context.Context, symbol string) (string, string, error) {
	resp, err := c.future.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return "", "", err
	}

	for _, s := range resp.Symbols {
		if s.Symbol == symbol {
			lotSizeFilter := s.LotSizeFilter()
			priceFilter := s.PriceFilter()
			return priceFilter.TickSize, lotSizeFilter.StepSize, nil
		}
	}

	return "", "", err
}

func (c *BinanceClient) LotSize(ctx context.Context, symbol string) (*futures.LotSizeFilter, error) {
	resp, err := c.future.NewExchangeInfoService().Do(ctx)
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

func (c *BinanceClient) Balance(ctx context.Context) (*futures.Balance, error) {
	resp, err := c.future.NewGetBalanceService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, b := range resp {
		if b.Asset == "USDT" {
			return b, nil
		}
	}

	return nil, nil

}

func (c *BinanceClient) Ticker(ctx context.Context, symbol string) (*futures.SymbolPrice, error) {
	resp, err := c.future.NewListPricesService().Symbol(symbol).Do(ctx)
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

func (c *BinanceClient) PositionWithoutSideIncludeEmpty(ctx context.Context, symbol string) (*futures.AccountPosition, error) {
	resp, err := c.future.NewGetAccountService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Positions {
		if p.Symbol == symbol {
			return p, nil
		}
	}

	return nil, nil
}

func (c *BinanceClient) PositionWithoutSide(ctx context.Context, symbol string) (*futures.AccountPosition, error) {
	resp, err := c.future.NewGetAccountService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Positions {
		if p.Symbol == symbol {
			positionAmt := str.New(p.PositionAmt).ToFloat64Default(0)
			if positionAmt == 0 {
				continue
			}

			return p, nil
		}
	}

	return nil, nil
}

func (c *BinanceClient) PositionSide(p futures.AccountPosition) (string, error) {
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

func (c *BinanceClient) PositionWithEmptyValue(ctx context.Context, symbol, holdSide string) (*futures.AccountPosition, error) {
	resp, err := c.future.NewGetAccountService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Positions {
		if p.Symbol == symbol {
			return p, nil
		}
	}

	return nil, nil
}

func (c *BinanceClient) Position(ctx context.Context, symbol, holdSide string) (*futures.AccountPosition, error) {
	resp, err := c.future.NewGetAccountService().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Positions {
		if p.Symbol == symbol {
			positionAmt := str.New(p.PositionAmt).ToFloat64Default(0)
			if positionAmt == 0 {
				continue
			}

			side := HoldSideShort
			if positionAmt > 0 {
				side = HoldSideLong
			}

			if side == holdSide {
				return p, nil
			}
		}
	}

	return nil, nil
}

func (c *BinanceClient) PlaceOrderAtPrice(ctx context.Context, symbol, holdSide, size, price string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeLimit).
		Side(c.convertHoldSide(holdSide)).
		Quantity(size).
		Price(price).
		TimeInForce(futures.TimeInForceTypeGTC).
		Do(ctx)
	if err != nil {
		return nil, c.errorConverter(err)
	}

	return resp, nil
}

func (c *BinanceClient) PlaceOrder(ctx context.Context, symbol, holdSide, size string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeMarket).
		Side(c.convertHoldSide(holdSide)).
		Quantity(size).
		PositionSide(futures.PositionSideTypeBoth).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceClient) SetTP(ctx context.Context, symbol, holdSide, price string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeTakeProfitMarket).
		Side(c.oppositeHoldSide(holdSide)).
		StopPrice(price).
		ClosePosition(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceClient) SetTPLimit(ctx context.Context, symbol, holdSide, price, size string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeTakeProfit).
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

func (c *BinanceClient) SetSL(ctx context.Context, symbol, holdSide, price string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeStopMarket).
		Side(c.oppositeHoldSide(holdSide)).
		StopPrice(price).
		ClosePosition(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceClient) SetSLLimit(ctx context.Context, symbol, holdSide, price, size string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeStop).
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

func (c *BinanceClient) SetTS(ctx context.Context, symbol, side, triggerPrice, callbackRate, size string) (*futures.CreateOrderResponse, error) {
	resp, err := c.future.NewCreateOrderService().
		Symbol(symbol).
		Type(futures.OrderTypeTrailingStopMarket).
		Side(c.oppositeHoldSide(side)).
		ActivationPrice(triggerPrice).
		CallbackRate(callbackRate).
		Quantity(size).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceClient) SetLeverage(ctx context.Context, symbol string, leverage int) (*futures.SymbolLeverage, error) {
	resp, err := c.future.NewChangeLeverageService().
		Symbol(symbol).
		Leverage(leverage).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *BinanceClient) SetMarginTypeToIsolate(ctx context.Context, symbol string) error {
	err := c.future.NewChangeMarginTypeService().
		Symbol(symbol).
		MarginType(futures.MarginTypeIsolated).
		Do(ctx)

	if err != nil {
		// Ignore error if margin type is already set to cross
		if err.Error() == "<APIError> code=-4046, msg=No need to change margin type." {
			return nil
		}
		return err
	}

	return nil
}

func (c *BinanceClient) SetMarginTypeToCross(ctx context.Context, symbol string) error {
	err := c.future.NewChangeMarginTypeService().
		Symbol(symbol).
		MarginType(futures.MarginTypeCrossed).
		Do(ctx)

	if err != nil {
		// Ignore error if margin type is already set to cross
		if err.Error() == "<APIError> code=-4046, msg=No need to change margin type." {
			return nil
		}
		return err
	}

	return nil
}

func (c *BinanceClient) OpenTPSLTSOrders(ctx context.Context, symbol string) ([]*futures.Order, error) {
	resp, err := c.future.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]*futures.Order, 0)
	for _, order := range resp {
		if order.Type == futures.OrderTypeStopMarket ||
			order.Type == futures.OrderTypeTakeProfitMarket ||
			order.Type == futures.OrderTypeTrailingStopMarket {
			orders = append(orders, order)
		}
	}

	return orders, nil
}

func (c *BinanceClient) OpenPositionOrders(ctx context.Context, symbol string) ([]*futures.Order, error) {
	resp, err := c.future.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]*futures.Order, 0)
	for _, order := range resp {
		if order.Type == futures.OrderTypeLimit ||
			order.Type == futures.OrderTypeMarket {
			orders = append(orders, order)
		}
	}

	return orders, nil
}

func (c *BinanceClient) CancelOpenOrders(ctx context.Context, symbol string, orderIDList []int64) ([]*futures.CancelOrderResponse, error) {
	cancelResult, err := c.future.NewCancelMultipleOrdersService().
		Symbol(symbol).
		OrderIDList(orderIDList).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return cancelResult, nil
}

func (c *BinanceClient) LimitOrder(ctx context.Context, symbol string) (*futures.Order, error) {
	resp, err := c.future.NewListOpenOrdersService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, order := range resp {
		if order.Type == futures.OrderTypeLimit &&
			order.WorkingType == futures.WorkingTypeContractPrice &&
			order.Symbol == symbol {
			return order, nil
		}
	}

	return nil, nil
}

func (c *BinanceClient) NotionAndLeverageBrackets(ctx context.Context, symbol string) (*futures.LeverageBracket, error) {
	resp, err := c.future.NewGetLeverageBracketService().
		Symbol(symbol).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, bracket := range resp {
		if bracket.Symbol == symbol {
			return bracket, nil
		}
	}

	return nil, nil
}

func (c *BinanceClient) oppositeHoldSide(holdSide string) futures.SideType {
	side := c.convertHoldSide(holdSide)
	switch side {
	case futures.SideTypeBuy:
		return futures.SideTypeSell
	case futures.SideTypeSell:
		return futures.SideTypeBuy
	default:
		return side
	}
}

func (c *BinanceClient) convertHoldSide(holdSide string) futures.SideType {
	switch holdSide {
	case HoldSideLong:
		return futures.SideTypeBuy
	case HoldSideShort:
		return futures.SideTypeSell
	default:
		return futures.SideType(holdSide)
	}
}

func (c *BinanceClient) errorConverter(apiErr error) error {
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

	return apiErr
}
