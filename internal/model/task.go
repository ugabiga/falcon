package model

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

var (
	ErrInvalidTaskType = errors.New("invalid_task_type")
)

const (
	TaskTypeLongGrid = "buying_grid"
	TaskTypeDCA      = "dca"
)

type Task struct {
	ID                string                 `json:"id" validate:"required"`
	UserID            string                 `json:"user_id" validate:"required"`
	TradingAccountID  string                 `json:"trading_account_id" validate:"required"`
	Currency          string                 `json:"currency" validate:"required"`
	Size              float64                `json:"size" validate:"required"`
	Symbol            string                 `json:"symbol" validate:"required"`
	Cron              string                 `json:"cron" validate:"required"`
	NextExecutionTime time.Time              `json:"next_execution_time" validate:"required"`
	IsActive          bool                   `json:"is_active" validate:"required"`
	Type              string                 `json:"type" validate:"required"`
	Params            map[string]interface{} `json:"params"`
	UpdatedAt         time.Time              `json:"updated_at" validate:"required"`
	CreatedAt         time.Time              `json:"created_at" validate:"required"`
}

func (t Task) GridParams() (*TaskGridParams, error) {
	if t.Type != TaskTypeLongGrid {
		return nil, ErrInvalidTaskType
	}

	marshalParams, err := json.Marshal(t.Params)
	if err != nil {
		return nil, err
	}

	var gridParams TaskGridParams
	if err := json.Unmarshal(marshalParams, &gridParams); err != nil {
		log.Printf("Error unmarshalling grid params. Err: %v", err)
		return nil, err
	}

	return &gridParams, nil
}
func (t Task) GridParamsV2() (*TaskGridParamsV2, error) {
	if t.Type != TaskTypeLongGrid {
		return nil, ErrInvalidTaskType
	}

	marshalParams, err := json.Marshal(t.Params)
	if err != nil {
		return nil, err
	}

	var gridParams TaskGridParamsV2
	if err := json.Unmarshal(marshalParams, &gridParams); err != nil {
		log.Printf("Error unmarshalling grid params. Err: %v", err)
		return nil, err
	}

	return &gridParams, nil
}

type TaskGridParams struct {
	GapPercent float64 `json:"gap_percent"`
	Quantity   int64   `json:"quantity"`
}

func (t TaskGridParams) ToParams() map[string]interface{} {
	return map[string]interface{}{
		"gap_percent": t.GapPercent,
		"quantity":    t.Quantity,
	}
}

type TaskGridParamsV2 struct {
	GapPercent           float64 `json:"gap_percent"`
	Quantity             int64   `json:"quantity"`
	UseIncrementalSize   bool    `json:"use_incremental_size"`
	IncrementalSize      float64 `json:"incremental_size"`
	DeletePreviousOrders bool    `json:"delete_previous_orders"`
}

func (t TaskGridParamsV2) ToParams() map[string]interface{} {
	return map[string]interface{}{
		"gap_percent":            t.GapPercent,
		"quantity":               t.Quantity,
		"use_incremental_size":   t.UseIncrementalSize,
		"incremental_size":       t.IncrementalSize,
		"delete_previous_orders": t.DeletePreviousOrders,
	}
}
