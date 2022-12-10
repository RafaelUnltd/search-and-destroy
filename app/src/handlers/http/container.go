package http

import (
	"search-and-destroy/app/src/handlers/http/players"
	"search-and-destroy/app/src/interfaces"
	"search-and-destroy/app/src/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	public := httpServer.Group("/public")
	public.POST("/login", h.PlayersHandler.AuthenticatePlayer)
	public.POST("/register", h.PlayersHandler.CreatePlayer)

	private := httpServer.Group("/private")
	private.Use(middleware.JWT([]byte("TODO")))
	private.GET("/players", h.PlayersHandler.ListPlayers)
}
