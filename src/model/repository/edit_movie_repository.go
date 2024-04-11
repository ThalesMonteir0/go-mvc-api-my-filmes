package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (mr *movieRepositoryInterface) UpdateMovie(id int, movieDomain model.MovieDomainInterface) *rest_err.RestErr {
	movieEntity := converter.ConvertMovieDomainToEntity(movieDomain)
	var updatedID int

	if err := mr.DB.QueryRow(SqlUpdateMovie, movieEntity.Name, movieEntity.Genre, movieEntity.Description, movieEntity.UrlImg, id).Scan(&updatedID); err != nil {
		return rest_err.NewInternalServerError("Unable to update movie")
	}

	return nil
}
