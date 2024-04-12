package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (mr *movieRepositoryInterface) CreateMovie(movieDomain model.MovieDomainInterface) *rest_err.RestErr {
	var movieIDInserted int
	movieEntity := converter.ConvertMovieDomainToEntity(movieDomain)

	err := mr.DB.QueryRow(SqlCreateMovie, movieEntity.Name, movieEntity.Genre, movieEntity.Description, movieEntity.UrlImg).Scan(&movieIDInserted)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
