package upbit

import "time"

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type CreateOrderResponse struct {
	UUID           string    `json:"uuid"`
	Side           string    `json:"side"`
	OrdType        string    `json:"ord_type"`
	Price          string    `json:"price"`
	State          string    `json:"state"`
	Market         string    `json:"market"`
	CreatedAt      time.Time `json:"created_at"`
	ReservedFee    string    `json:"reserved_fee"`
	RemainingFee   string    `json:"remaining_fee"`
	PaidFee        string    `json:"paid_fee"`
	Locked         string    `json:"locked"`
	ExecutedVolume string    `json:"executed_volume"`
	TradesCount    int       `json:"trades_count"`
}

type Ticker struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKst       string  `json:"trade_date_kst"`
	TradeTimeKst       string  `json:"trade_time_kst"`
	TradeTimestamp     int64   `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24H   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24H  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

type OrderChange struct {
	BidFee      string `json:"bid_fee"`
	AskFee      string `json:"ask_fee"`
	MakerBidFee string `json:"maker_bid_fee"`
	MakerAskFee string `json:"maker_ask_fee"`
	Market      struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		OrderSides []string `json:"order_sides"`
		BidTypes   []string `json:"bid_types"`
		AskTypes   []string `json:"ask_types"`
		Bid        struct {
			Currency string `json:"currency"`
			MinTotal string `json:"min_total"`
		} `json:"bid"`
		Ask struct {
			Currency string `json:"currency"`
			MinTotal string `json:"min_total"`
		} `json:"ask"`
		MaxTotal string `json:"max_total"`
		State    string `json:"state"`
	} `json:"market"`
	BidAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"bid_account"`
	AskAccount struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"ask_account"`
}

type OrderBooks []OrderBook

type OrderBook struct {
	Market         string `json:"market"`
	OrderbookUnits []struct {
		AskPrice int64   `json:"ask_price"`
		AskSize  float64 `json:"ask_size"`
		BidPrice int64   `json:"bid_price"`
		BidSize  float64 `json:"bid_size"`
	} `json:"orderbook_units"`
	Timestamp    int64   `json:"timestamp"`
	TotalAskSize float64 `json:"total_ask_size"`
	TotalBidSize float64 `json:"total_bid_size"`
}

func (o OrderBook) UnitPrice() float64 {
	unitPrice := int64(0)
	for i, orderBookUnit := range o.OrderbookUnits {
		if i == 0 {
			continue
		}
		previousOrderBookUnit := o.OrderbookUnits[i-1]
		currentOrderBookUnit := orderBookUnit
		subtract := currentOrderBookUnit.AskPrice - previousOrderBookUnit.AskPrice

		if unitPrice == 0 {
			unitPrice = subtract
		}
		if subtract < unitPrice {
			unitPrice = subtract
		}
	}

	return float64(unitPrice)
}

type Order struct {
	UUID            string    `json:"uuid"`
	Side            string    `json:"side"`
	OrdType         string    `json:"ord_type"`
	Price           string    `json:"price"`
	State           string    `json:"state"`
	Market          string    `json:"market"`
	CreatedAt       time.Time `json:"created_at"`
	Volume          string    `json:"volume"`
	RemainingVolume string    `json:"remaining_volume"`
	ReservedFee     string    `json:"reserved_fee"`
	RemainingFee    string    `json:"remaining_fee"`
	PaidFee         string    `json:"paid_fee"`
	Locked          string    `json:"locked"`
	ExecutedVolume  string    `json:"executed_volume"`
	TradesCount     int       `json:"trades_count"`
}
