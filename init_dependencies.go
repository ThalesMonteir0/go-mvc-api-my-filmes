package main

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/lists"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/movie"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/service"
)

func initDependencies(db *sql.DB) (user.UserControllerInterface, movie.MovieControllerInterface, lists.ListControllerInterface) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserDomainService(userRepository)

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)

	listRepository := repository.NewListRepository(db)
	listService := service.NewListService(listRepository)

	return user.NewUserController(userService), movie.NewMovieController(movieService), lists.NewListController(listService)
}
