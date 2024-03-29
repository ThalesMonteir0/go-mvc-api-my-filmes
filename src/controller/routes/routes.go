package routes

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/gofiber/fiber/v2"
)

func InitRoutesV1(v1 fiber.Router, userController user.UserControllerInterface) {
	v1.Get("/user/:id", userController.FindUserByID)
	v1.Post("/user", userController.CreateUser)
	v1.Delete("/user/:id", userController.DeleteUser)
	v1.Put("/user/:id", userController.EditUser)

}
