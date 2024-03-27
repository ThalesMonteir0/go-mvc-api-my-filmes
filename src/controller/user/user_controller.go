package user

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService service.UserDomainService) UserControllerInterface {
	return &userControllerService{
		service: userService,
	}
}

type UserControllerInterface interface {
	CreateUser(c *fiber.Ctx) error
	FindUserByID(c *fiber.Ctx) error
	EditUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userControllerService struct {
	service service.UserDomainService
}
