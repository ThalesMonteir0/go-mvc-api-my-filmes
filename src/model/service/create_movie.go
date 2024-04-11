package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ms *movieServiceInterface) CreateMovie(movie model.MovieDomainInterface) *rest_err.RestErr {
	if movieDomain, err := ms.findMovieByName(movie.GetName()); movieDomain.GetID() != 0 || err == nil {
		return rest_err.NewBadRequestError("movie exists!")
	}

	movie.TransformNameToUpperCase()
	movie.TransformDescriptionToUpperCase()
	movie.TransformGenreToUpperCase()

	if err := ms.repository.CreateMovie(movie); err != nil {
		return err
	}

	return nil
}
