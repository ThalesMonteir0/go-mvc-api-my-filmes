package repository

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

type movieRepositoryInterface struct {
	DB *sql.DB
}

type MovieRepositoryInterface interface {
	FindMovieByID(int) (*model.MovieDomainInterface, *rest_err.RestErr)
	CreateMovie(model.MovieDomainInterface) *rest_err.RestErr
	DeleteMovie(int) *rest_err.RestErr
	UpdateMovie(int, model.MovieDomainInterface) *rest_err.RestErr
	FindMovieByName(name string) (*model.MovieDomainInterface, *rest_err.RestErr)
}

func NewMovieRepository(db *sql.DB) MovieRepositoryInterface {
	return &movieRepositoryInterface{
		DB: db,
	}

}
