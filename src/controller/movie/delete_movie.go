package movie

import "github.com/gofiber/fiber/v2"

func (mc *movieControllerInterface) DeleteMovie(c *fiber.Ctx) error {
	movieID, _ := c.ParamsInt("id")

	if err := mc.service.DeleteMovie(movieID); err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON("Movie Deleted!")
}
