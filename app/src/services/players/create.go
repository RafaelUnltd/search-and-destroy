package players

import (
	"log"
	"poly-shooters/app/src/entities"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) CreatePlayer(player entities.Player) (entities.Player, error) {
	// Generates hashed password to store in database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Couldn't generate hashed password for string: ", err)
		return entities.Player{}, err
	}

	return s.PlayersRepository.CreatePlayer(
		player.Username,
		player.Email,
		string(hashedPassword[:]),
	)
}
