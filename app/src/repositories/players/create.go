package players

import (
	"poly-shooters/app/src/entities"
	"time"
)

func (r Repository) CreatePlayer(username, email, password string) (entities.Player, error) {
	player := entities.Player{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := r.Database.Create(&player)
	if result.Error != nil {
		return entities.Player{}, result.Error
	}

	return player, nil
}
