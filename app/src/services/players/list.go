package players

import "poly-shooters/app/src/entities"

func (s Service) ListPlayers() []entities.Player {
	return s.PlayersRepository.ListPlayers()
}
