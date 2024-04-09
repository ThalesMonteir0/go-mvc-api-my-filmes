package movie

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func (mc *movieControllerInterface) CreateMovie(c *fiber.Ctx) error {
	var movieRequest request.MovieRequest

	errParse := c.BodyParser(&movieRequest)
	if errParse != nil {
		err := rest_err.NewBadRequestError("Invalid Parameters")
		return c.Status(err.Code).JSON(err.Message)
	}

	movieDomain := model.NewMovieDomain(movieRequest.Name, movieRequest.Genre, movieRequest.Description, movieRequest.UrlImg)
	err := mc.service.CreateMovie(movieDomain)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Status(fiber.StatusOK).JSON("movie created")
}
