package movie

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/view"
	"github.com/gofiber/fiber/v2"
)

func (mc *movieControllerInterface) FindMovieByID(c *fiber.Ctx) error {
	movieID, _ := c.ParamsInt("id")

	movieDomain, err := mc.service.GetMovieByID(movieID)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON(view.ConvertMovieDomainToResponse(movieDomain))
}
