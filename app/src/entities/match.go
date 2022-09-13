package entities

import (
	"time"

	"gorm.io/gorm"
)

// Match is a match entity representation
type Match struct {
	ID         string         `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Status     string         `json:"status"`
	WinnerTeam string         `json:"winner_team"`
	CreatedAt  time.Time      `json:"created_at" validate:"omitempty"`
	UpdatedAt  time.Time      `json:"update_at" validate:"omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" validate:"omitempty"`
}
