package models

import (
	"time"

	"github.com/google/uuid"
)

// The `json:"xxx"` tags are used for JSON serialization
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}
