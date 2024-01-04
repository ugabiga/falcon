package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
