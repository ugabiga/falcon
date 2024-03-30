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
	ID        string    `json:"id,omitempty" validate:"required"`
	UserID    string    `json:"user_id,omitempty" validate:"required"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Exchange  string    `json:"exchange,omitempty" validate:"required"`
	IP        string    `json:"ip,omitempty" validate:"required"`
	Key       string    `json:"key,omitempty"  validate:"required"`
	Secret    string    `json:"secret,omitempty"`
	Phrase    string    `json:"phrase,omitempty" validate:"required"`
	UpdatedAt time.Time `json:"updated_at,omitempty" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" validate:"required"`
}

type TaskHistory struct {
	ID               string    `json:"id" validate:"required"`
	TaskID           string    `json:"task_id" validate:"required"`
	TradingAccountID string    `json:"trading_account_id" validate:"required"`
	UserID           string    `json:"user_id" validate:"required"`
	IsSuccess        bool      `json:"is_success" validate:"required"`
	Log              string    `json:"log" validate:"required"`
	UpdatedAt        time.Time `json:"updated_at" validate:"required"`
	CreatedAt        time.Time `json:"created_at" validate:"required"`
}

type StaticIP struct {
	ID             string    `json:"id"`
	IPAddress      string    `json:"ip_address"`
	IPAvailability bool      `json:"ip_availability"`
	IPUsageCount   int       `json:"ip_usage_count"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}
