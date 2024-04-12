package service

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

func (ms *movieServiceInterface) DeleteMovie(movieID int) *rest_err.RestErr {
	if movieExists, _ := ms.GetMovieByID(movieID); movieExists == nil {
		return rest_err.NewBadRequestError("Movie not exists")
	}

	if err := ms.repository.DeleteMovie(movieID); err != nil {
		return err
	}

	return nil
}
