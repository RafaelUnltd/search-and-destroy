package players

import (
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewPlayersRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}
