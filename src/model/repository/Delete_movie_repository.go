package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
)

func (mr *movieRepositoryInterface) DeleteMovie(movieID int) *rest_err.RestErr {
	err := mr.DB.QueryRow(SqlDeleteMovie, movieID).Err()
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
