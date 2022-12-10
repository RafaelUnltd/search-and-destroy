package main

import (
	"log"
	"search-and-destroy/app/cmd/start"
	"search-and-destroy/app/src/repositories"
	"search-and-destroy/app/src/services"
)

func main() {
	// Starts dabatase connection
	database, err := start.StartDatabase()
	if err != nil {
		log.Fatalf("Couldn't connect to database, with error: %s. Ending proccess.", err.Error())
	}

	// Desclares repositories and services for all the apis
	repositoriesContainer := repositories.NewRepositoriesContainer(database)
	servicesContainer := services.NewServicesContainer(repositoriesContainer)

	// Starts http server and register it's routes
	httpServer, err := start.StartHttpServer(database, servicesContainer)
	if err != nil {
		log.Fatalf("Couldn't start http service, with error: %s. Ending proccess.", err.Error())
	}

	httpServer.Logger.Fatal(httpServer.Start(":4001"))
}
