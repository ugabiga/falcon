package model

import (
	"time"
)

const (
	ExchangeBinanceFutures = "binance_futures"
	ExchangeBinanceSpot    = "binance_spot"
	ExchangeUpbit          = "upbit"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Authentication struct {
	ID         string    `json:"id"` //provider + identifier
	UserID     string    `json:"user_id"`
	Provider   string    `json:"provider"`
	Identifier string    `json:"identifier"`
	Credential string    `json:"credential"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type TradingAccount struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Exchange  string    `json:"exchange,omitempty"`
	IP        string    `json:"ip,omitempty"`
	Key       string    `json:"key,omitempty"`
	Secret    string    `json:"secret,omitempty"`
	Phrase    string    `json:"phrase,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TaskHistory struct {
	ID               string    `json:"id"`
	TaskID           string    `json:"task_id"`
	TradingAccountID string    `json:"trading_account_id"`
	UserID           string    `json:"user_id"`
	IsSuccess        bool      `json:"is_success"`
	Log              string    `json:"log"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
}

type StaticIP struct {
	ID             string    `json:"id"`
	IPAddress      string    `json:"ip_address"`
	IPAvailability bool      `json:"ip_availability"`
	IPUsageCount   int       `json:"ip_usage_count"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}
