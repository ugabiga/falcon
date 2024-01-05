package model

import "time"

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
