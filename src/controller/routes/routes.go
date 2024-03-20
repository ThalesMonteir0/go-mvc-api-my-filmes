package routes

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/user"
	"github.com/gofiber/fiber/v2"
)

func InitRoutesV1(v1 fiber.Router) {
	v1.Get("/getUserByID", user.FindUserByID)
	v1.Post("/createUser", user.CreateUser)

}
