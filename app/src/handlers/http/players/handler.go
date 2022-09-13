package players

import (
	"poly-shooters/app/src/interfaces"
	"poly-shooters/app/src/services"
)

type Handler struct {
	PlayersService interfaces.PlayersServiceInterface
}

func NewPlayersHandler(s services.ServicesContainer) interfaces.PlayersHandlerInterface {
	return Handler{
		PlayersService: s.PlayersService,
	}
}
