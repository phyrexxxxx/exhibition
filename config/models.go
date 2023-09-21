package config

import (
	"time"

	"github.com/google/uuid"
	"github.com/phyrexxxxx/exhibition/internal/database"
)

// The `json:"xxx"` tags are used for JSON serialization
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func DBUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
	}
}
