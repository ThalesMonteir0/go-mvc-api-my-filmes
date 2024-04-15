package routes

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/lists"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/movie"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func InitRoutesV1(v1 fiber.Router, userController user.UserControllerInterface, movieController movie.MovieControllerInterface, listController lists.ListControllerInterface) {
	v1.Get("/user/:id", model.VerifyTokenMiddleware, userController.FindUserByID)
	v1.Post("/user", model.VerifyTokenMiddleware, userController.CreateUser)
	v1.Delete("/user/:id", model.VerifyTokenMiddleware, userController.DeleteUser)
	v1.Put("/user/:id", model.VerifyTokenMiddleware, userController.EditUser)
	v1.Post("/login", userController.LoginUser)

	v1.Get("/movie/:id", model.VerifyTokenMiddleware, movieController.FindMovieByID)
	v1.Post("/movie", model.VerifyTokenMiddleware, movieController.CreateMovie)
	v1.Put("/movie/:id", model.VerifyTokenMiddleware, movieController.UpdateMovie)
	v1.Delete("/movie/:id", model.VerifyTokenMiddleware, movieController.DeleteMovie)

	v1.Get("/list/:userID", model.VerifyTokenMiddleware, listController.GetListsByUserID)
	v1.Post("/list", model.VerifyTokenMiddleware, listController.CreateList)
	v1.Delete("/list", model.VerifyTokenMiddleware, listController.DeleteList)
}
