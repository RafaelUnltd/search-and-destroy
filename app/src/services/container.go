package services

import (
	"poly-shooters/app/src/handlers/http/interfaces"
	"poly-shooters/app/src/repositories"
	"poly-shooters/app/src/services/players"
)

type ServicesContainer struct {
	PlayersService interfaces.PlayersServiceInterface
}

func NewServicesContainer(repositories repositories.RepositoriesContainer) ServicesContainer {
	return ServicesContainer{
		PlayersService: players.NewPlayersService(repositories),
	}
}
