package players

import (
	"search-and-destroy/app/src/entities"
	"search-and-destroy/app/src/structs"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) AuthenticatePlayer(playerData entities.Player) (entities.Player, string, error) {
	player := s.PlayersRepository.GetOnePlayer(entities.Player{Username: playerData.Username})

	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(playerData.Password)); err != nil {
		return entities.Player{}, "", err
	}

	claims := &structs.AuthenticationClaims{
		ID:       player.ID,
		Username: player.Username,
		Email:    player.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte("TODO"))
	if err != nil {
		return entities.Player{}, "", err
	}

	return player, token, nil
}
