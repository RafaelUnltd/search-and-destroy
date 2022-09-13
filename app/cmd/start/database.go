package start

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDatabase() (*gorm.DB, error) {
	dsn := "host=database-poly-shooters user=root password=root dbname=shooters port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Databse successfully connected!")

	return database, nil
}
