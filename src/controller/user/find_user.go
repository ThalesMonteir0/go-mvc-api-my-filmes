package user

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/view"
	"github.com/gofiber/fiber/v2"
)

func (uc *userControllerService) FindUserByID(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("id")

	userDomain, err := uc.service.FindUserByID(userID)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}
	userResponse := view.ConvertUserDomainToResponse(userDomain)
	return c.Status(fiber.StatusOK).JSON(userResponse)
}
