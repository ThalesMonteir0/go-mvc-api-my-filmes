package routes

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func InitRoutesV1(v1 fiber.Router, userController user.UserControllerInterface) {
	v1.Get("/user/:id", model.VerifyTokenMiddleware, userController.FindUserByID)
	v1.Post("/user", model.VerifyTokenMiddleware, userController.CreateUser)
	v1.Delete("/user/:id", model.VerifyTokenMiddleware, userController.DeleteUser)
	v1.Put("/user/:id", model.VerifyTokenMiddleware, userController.EditUser)
	v1.Post("/login", userController.LoginUser)
}
