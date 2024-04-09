package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ms *movieServiceInterface) GetMovieByID(id int) (model.MovieDomainInterface, *rest_err.RestErr) {
	movieDomain, err := ms.repository.FindMovieByID(id)
	if err != nil {
		return nil, err
	}

	return *movieDomain, nil
}

func (ms *movieServiceInterface) findMovieByName(name string) (model.MovieDomainInterface, *rest_err.RestErr) {
	movieDomain, err := ms.repository.FindMovieByName(name)
	if err != nil {
		return nil, err
	}

	return *movieDomain, nil
}
