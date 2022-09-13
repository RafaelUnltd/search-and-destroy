package players

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) ListPlayers(c echo.Context) error {
	players := h.PlayersService.ListPlayers()
	return c.JSON(http.StatusOK, players)
}
