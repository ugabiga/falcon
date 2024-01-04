package model

import "time"

type Authentication struct {
	ID         string    `json:"id"` //provider + identifier
	UserID     string    `json:"user_id"`
	Provider   string    `json:"provider"`
	Identifier string    `json:"identifier"`
	Credential string    `json:"credential"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}
