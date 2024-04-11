package movie

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func (mc *movieControllerInterface) UpdateMovie(c *fiber.Ctx) error {
	var movieRequest request.MovieRequest
	movieID, _ := c.ParamsInt("id")

	if err := c.BodyParser(&movieRequest); err != nil {
		badRequest := rest_err.NewBadRequestError("Invalid Parameters")
		return c.Status(badRequest.Code).JSON(badRequest.Message)
	}

	movieDomain := model.NewMovieDomain(movieRequest.Name, movieRequest.Genre, movieRequest.Description, movieRequest.UrlImg)
	if err := mc.service.UpdateMovie(movieID, movieDomain); err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON("updated")
}
