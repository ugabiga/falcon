package model

import "time"

type Task struct {
	ID                string                 `json:"id"`
	UserID            string                 `json:"user_id"`
	TradingAccountID  string                 `json:"trading_account_id"`
	Currency          string                 `json:"currency"`
	Size              float64                `json:"size"`
	Symbol            string                 `json:"symbol"`
	Cron              string                 `json:"cron"`
	NextExecutionTime time.Time              `json:"next_execution_time"`
	IsActive          bool                   `json:"is_active"`
	Type              string                 `json:"type"`
	Params            map[string]interface{} `json:"params"`
	UpdatedAt         time.Time              `json:"updated_at"`
	CreatedAt         time.Time              `json:"created_at"`
}
