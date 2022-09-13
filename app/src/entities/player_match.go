package entities

import (
	"time"

	"gorm.io/gorm"
)

// PlayerMatch is a player_match entity representation
type PlayerMatch struct {
	ID        string         `json:"id" gorm:"primary_key"`
	PlayerID  string         `json:"player_id"`
	MatchID   string         `json:"match_id"`
	Team      string         `json:"team"`
	Kills     int            `json:"kill"`
	Deaths    int            `json:"deaths"`
	CreatedAt time.Time      `json:"created_at" validate:"omitempty"`
	UpdatedAt time.Time      `json:"update_at" validate:"omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" validate:"omitempty"`
}
