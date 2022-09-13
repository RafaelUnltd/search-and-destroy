package repositories

import (
	"poly-shooters/app/src/interfaces"
	"poly-shooters/app/src/repositories/players"

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
