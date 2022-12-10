package players

import "search-and-destroy/app/src/entities"

func (r Repository) ListPlayers() []entities.Player {
	var players []entities.Player
	r.Database.Find(&players)
	return players
}
