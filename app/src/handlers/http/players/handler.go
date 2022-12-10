package players

import (
	"search-and-destroy/app/src/interfaces"
	"search-and-destroy/app/src/services"
)

type Handler struct {
	PlayersService interfaces.PlayersServiceInterface
}

func NewPlayersHandler(s services.ServicesContainer) interfaces.PlayersHandlerInterface {
	return Handler{
		PlayersService: s.PlayersService,
	}
}
