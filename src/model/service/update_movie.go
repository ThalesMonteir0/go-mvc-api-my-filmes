package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ms *movieServiceInterface) UpdateMovie(id int, movie model.MovieDomainInterface) *rest_err.RestErr {
	if movieDomain, _ := ms.repository.FindMovieByID(id); movieDomain == nil {
		return rest_err.NewBadRequestError("Movie not exists")
	}

	movie.TransformNameToUpperCase()
	movie.TransformDescriptionToUpperCase()
	movie.TransformGenreToUpperCase()

	if err := ms.repository.UpdateMovie(id, movie); err != nil {
		return err
	}

	return nil
}
