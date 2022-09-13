package players

import "poly-shooters/app/src/entities"

func (r Repository) GetOne(playerSearch entities.Player) entities.Player {
	var player entities.Player
	r.Database.Model(playerSearch).First(&player)
	return player
}
