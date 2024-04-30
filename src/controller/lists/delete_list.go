package lists

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/gofiber/fiber/v2"
)

func (lc *listController) DeleteList(c *fiber.Ctx) error {
	listID, err := c.ParamsInt("listID")
	if err != nil || listID == 0 {
		errParams := rest_err.NewBadRequestError("params invalid")
		return c.Status(errParams.Code).JSON(errParams.Message)
	}

	if err := lc.service.DeleteList(listID); err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON("Deleted")
}
