package interfaces

import (
	"poly-shooters/app/src/entities"

	"github.com/labstack/echo/v4"
)

type PlayersHandlerInterface interface {
	ListPlayers(c echo.Context) error
	CreatePlayer(c echo.Context) error
}

type PlayersServiceInterface interface {
	ListPlayers() []entities.Player
	CreatePlayer(player entities.Player) (entities.Player, error)
}

type PlayersRepositoryInterface interface {
	ListPlayers() []entities.Player
	CreatePlayer(username, email, password string) (entities.Player, error)
	GetOne(playerSearch entities.Player) entities.Player
}
