package model

import "time"

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
