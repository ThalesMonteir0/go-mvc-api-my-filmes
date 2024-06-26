package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (mr *movieRepositoryInterface) FindMovieByID(id int) (*model.MovieDomainInterface, *rest_err.RestErr) {
	var movie entity.MovieEntity

	err := mr.DB.QueryRow(SqlFindMovieByID, id).Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Description, &movie.UrlImg)
	if err != nil {
		return nil, rest_err.NewNotFoundError("Movie not found")
	}

	return converter.ConvertMovieEntityToDomain(movie), nil
}

func (mr *movieRepositoryInterface) FindMovieByName(name string) (*model.MovieDomainInterface, *rest_err.RestErr) {
	var movieEntity entity.MovieEntity

	err := mr.DB.QueryRow(SqlFindMovieByName, name).Scan(&movieEntity.ID, &movieEntity.Name, &movieEntity.Genre, &movieEntity.Description, &movieEntity.UrlImg)
	if err != nil {
		return nil, rest_err.NewNotFoundError("movie name not found")
	}

	return converter.ConvertMovieEntityToDomain(movieEntity), nil
}
