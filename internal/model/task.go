package model

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrInvalidTaskType = errors.New("invalid_task_type")
)

const (
	TaskTypeGrid = "Grid"
	TaskTypeDCA  = "DCA"
)

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

type TaskGridParams struct {
	Gap      int64 `json:"gap"`
	Quantity int64 `json:"quantity"`
}

func (t Task) GridParams() (*TaskGridParams, error) {
	if t.Type != TaskTypeGrid {
		return nil, ErrInvalidTaskType
	}

	marshalParams, err := json.Marshal(t.Params)
	if err != nil {
		return nil, err
	}

	var gridParams TaskGridParams
	if err := json.Unmarshal(marshalParams, &gridParams); err != nil {
		return nil, err
	}

	return &gridParams, nil
}
