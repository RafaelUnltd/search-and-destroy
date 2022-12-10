package services

import (
	"search-and-destroy/app/src/interfaces"
	"search-and-destroy/app/src/repositories"
	"search-and-destroy/app/src/services/players"
)

type ServicesContainer struct {
	PlayersService interfaces.PlayersServiceInterface
}

func NewServicesContainer(repositories repositories.RepositoriesContainer) ServicesContainer {
	return ServicesContainer{
		PlayersService: players.NewPlayersService(repositories),
	}
}
