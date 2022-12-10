package interfaces

import (
	"search-and-destroy/app/src/entities"

	"github.com/labstack/echo/v4"
)

type PlayersHandlerInterface interface {
	ListPlayers(c echo.Context) error
	CreatePlayer(c echo.Context) error
	AuthenticatePlayer(c echo.Context) error
}

type PlayersServiceInterface interface {
	ListPlayers() []entities.Player
	CreatePlayer(player entities.Player) (entities.Player, error)
	AuthenticatePlayer(player entities.Player) (entities.Player, string, error)
}

type PlayersRepositoryInterface interface {
	ListPlayers() []entities.Player
	CreatePlayer(username, email, password string) (entities.Player, error)
	GetOnePlayer(playerSearch entities.Player) entities.Player
}
