package players

import (
	"net/http"
	"poly-shooters/app/src/entities"
	"poly-shooters/app/src/structs"

	"github.com/labstack/echo/v4"
)

func (h Handler) AuthenticatePlayer(c echo.Context) error {
	var playerData entities.Player

	err := c.Bind(&playerData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't read all data from player. Send again with username, password and email.")
	} else if playerData.Username == "" || playerData.Password == "" {
		return c.JSON(http.StatusBadRequest, "Please sent username and password in order to authenticate.")
	}

	player, token, err := h.PlayersService.AuthenticatePlayer(playerData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Incorrect username or password")
	}

	return c.JSON(http.StatusOK, structs.AuthenticatePlayerResponse{
		ID:       player.ID,
		Email:    player.Email,
		Username: player.Username,
		Token:    token,
	})
}
