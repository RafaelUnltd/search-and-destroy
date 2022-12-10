package players

import (
	"search-and-destroy/app/src/interfaces"
	"search-and-destroy/app/src/repositories"
)

type Service struct {
	PlayersRepository interfaces.PlayersRepositoryInterface
}

func NewPlayersService(repositories repositories.RepositoriesContainer) Service {
	return Service{
		PlayersRepository: repositories.PlayersRepository,
	}
}
