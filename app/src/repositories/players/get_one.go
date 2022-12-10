package players

import "search-and-destroy/app/src/entities"

func (r Repository) GetOnePlayer(playerSearch entities.Player) entities.Player {
	var player entities.Player
	r.Database.Model(playerSearch).First(&player)
	return player
}
