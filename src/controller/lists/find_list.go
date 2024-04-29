package lists

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/gofiber/fiber/v2"
)

func (lc *listController) GetListsByUserID(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("userID")
	if userID == 0 {
		err := rest_err.NewBadRequestError("UserID not found")
		return c.Status(err.Code).JSON(err.Message)
	}

	lists, err := lc.service.FindListsByUserID(userID)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON(lists)
}
