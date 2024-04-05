package movie

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/service"
	"github.com/gofiber/fiber/v2"
)

type movieControllerInterface struct {
	service service.MovieServiceInterface
}

type MovieControllerInterface interface {
	CreateMovie(*fiber.Ctx) error
	DeleteMovie(*fiber.Ctx) error
	UpdateMovie(*fiber.Ctx) error
	FindMovieByID(*fiber.Ctx) error
}

func NewMovieController(service service.MovieServiceInterface) MovieControllerInterface {
	return &movieControllerInterface{
		service: service,
	}
}
