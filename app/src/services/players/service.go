package players

import (
	"poly-shooters/app/src/interfaces"
	"poly-shooters/app/src/repositories"
)

type Service struct {
	PlayersRepository interfaces.PlayersRepositoryInterface
}

func NewPlayersService(repositories repositories.RepositoriesContainer) Service {
	return Service{
		PlayersRepository: repositories.PlayersRepository,
	}
}
