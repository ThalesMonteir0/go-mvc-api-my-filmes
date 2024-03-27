package main

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/database/postgresql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB, err := postgresql.NewPostgresConnection()
	if err != nil {

	}
	//init dependecias
	userController := initDependencies(DB)

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.InitRoutesV1(v1, userController)

	err = app.Listen(":5000")
	if err != nil {
		log.Fatal("Nâo foi possivel subir a aplicação!")
	}

}
