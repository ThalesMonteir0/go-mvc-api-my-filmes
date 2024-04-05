package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (mr *movieRepositoryInterface) UpdateMovie(id int, movieDomain model.MovieDomainInterface) *rest_err.RestErr {
	return nil
}
