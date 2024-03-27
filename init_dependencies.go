package main

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/service"
)

func initDependencies(db *sql.DB) user.UserControllerInterface {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserDomainService(userRepository)
	return user.NewUserController(userService)
}
