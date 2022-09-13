package entities

import (
	"time"

	"gorm.io/gorm"
)

// Player is a player entity representation
type Player struct {
	ID        string         `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at" validate:"omitempty"`
	UpdatedAt time.Time      `json:"update_at" validate:"omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" validate:"omitempty"`
}
