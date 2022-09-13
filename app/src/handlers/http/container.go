package http

import (
	"poly-shooters/app/src/handlers/http/interfaces"
	"poly-shooters/app/src/handlers/http/players"
	"poly-shooters/app/src/services"

	"github.com/labstack/echo/v4"
)

type HandlersContainer struct {
	PlayersHandler interfaces.PlayersHandlerInterface
}

func NewHandlersContainer(s services.ServicesContainer) HandlersContainer {
	return HandlersContainer{
		PlayersHandler: players.NewPlayersHandler(s),
	}
}

func (h HandlersContainer) RegisterRoutes(httpServer *echo.Echo) {
	players := httpServer.Group("/players")
	players.GET("", h.PlayersHandler.ListPlayers)
}
