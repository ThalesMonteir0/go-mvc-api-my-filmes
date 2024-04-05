package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository"
)

type movieServiceInterface struct {
	repository repository.MovieRepositoryInterface
}

type MovieServiceInterface interface {
	GetMovieByID(int) (model.MovieDomainInterface, *rest_err.RestErr)
	CreateMovie(model.MovieDomainInterface) *rest_err.RestErr
	DeleteMovie(int) *rest_err.RestErr
	UpdateMovie(int, model.MovieDomainInterface) *rest_err.RestErr
}

func NewMovieService(repository repository.MovieRepositoryInterface) MovieServiceInterface {
	return &movieServiceInterface{
		repository: repository,
	}
}
