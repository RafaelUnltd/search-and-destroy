package players

import "search-and-destroy/app/src/entities"

func (s Service) ListPlayers() []entities.Player {
	return s.PlayersRepository.ListPlayers()
}
