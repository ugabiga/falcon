package model

type Authentication struct {
	ID         string `json:"id"` //provider + identifier
	UserID     string `json:"user_id"`
	Provider   string `json:"provider"`
	Identifier string `json:"identifier"`
	Credential string `json:"credential"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
	User       *User  `json:"user,omitempty"`
}
