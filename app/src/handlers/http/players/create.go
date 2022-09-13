package players

import (
	"net/http"
	"poly-shooters/app/src/entities"

	"github.com/labstack/echo/v4"
)

func (h Handler) CreatePlayer(c echo.Context) error {
	var playerData entities.Player

	err := c.Bind(&playerData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't read all data from player. Send again with username, password and email.")
	}

	created, err := h.PlayersService.CreatePlayer(playerData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Couldn't create player.")
	}

	return c.JSON(http.StatusOK, created)
}
