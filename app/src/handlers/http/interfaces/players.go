package interfaces

import (
	"poly-shooters/app/src/entities"

	"github.com/labstack/echo/v4"
)

type PlayersHandlerInterface interface {
	ListPlayers(c echo.Context) error
}

type PlayersRepositoryInterface interface {
	ListPlayers() []entities.Player
}

type PlayersServiceInterface interface {
	ListPlayers() []entities.Player
}
