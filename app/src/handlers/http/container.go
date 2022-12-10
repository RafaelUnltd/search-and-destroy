package http

import (
	"poly-shooters/app/src/handlers/http/players"
	"poly-shooters/app/src/interfaces"
	"poly-shooters/app/src/services"

	"github.com/labstack/echo/middleware"
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
	httpServer.POST("/login", h.PlayersHandler.AuthenticatePlayer)

	players := httpServer.Group("/players")
	players.Use(middleware.JWT([]byte("TODO")))
	players.GET("", h.PlayersHandler.ListPlayers)
	players.POST("", h.PlayersHandler.CreatePlayer)
}
