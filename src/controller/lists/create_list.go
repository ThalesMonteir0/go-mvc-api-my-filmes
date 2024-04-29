package lists

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func (lc *listController) CreateList(c *fiber.Ctx) error {
	var listBody request.ListRequest

	errParse := c.BodyParser(&listBody)
	if errParse != nil {
		errRequest := rest_err.NewBadRequestError("struct error")
		return c.Status(errRequest.Code).JSON(errRequest.Message)
	}

	listDomain := model.NewListDomain(listBody.Title, listBody.Description, listBody.UserID)
	err := lc.service.CreateList(listDomain)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON("Created")
}
