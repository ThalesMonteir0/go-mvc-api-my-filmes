package routes

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/gofiber/fiber/v2"
)

func InitRoutesV1(v1 fiber.Router, userController user.UserControllerInterface) {
	v1.Get("/getUserByID", userController.FindUserByID)
	v1.Post("/createUser", userController.CreateUser)

}
