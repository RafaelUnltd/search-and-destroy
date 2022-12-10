package repositories

import (
	"search-and-destroy/app/src/interfaces"
	"search-and-destroy/app/src/repositories/players"

	"gorm.io/gorm"
)

type RepositoriesContainer struct {
	PlayersRepository interfaces.PlayersRepositoryInterface
}

func NewRepositoriesContainer(db *gorm.DB) RepositoriesContainer {
	return RepositoriesContainer{
		PlayersRepository: players.NewPlayersRepository(db),
	}
}
