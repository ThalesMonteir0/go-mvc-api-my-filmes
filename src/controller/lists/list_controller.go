package lists

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/service"
	"github.com/gofiber/fiber/v2"
)

type listController struct {
	service service.ListServiceInterface
}

type ListControllerInterface interface {
	GetListsByUserID(c *fiber.Ctx) error
	CreateList(c *fiber.Ctx) error
	DeleteList(c *fiber.Ctx) error
}

func NewListController(service service.ListServiceInterface) ListControllerInterface {
	return &listController{
		service: service,
	}
}
