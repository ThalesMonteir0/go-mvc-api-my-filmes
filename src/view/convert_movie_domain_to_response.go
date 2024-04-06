package view

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/response"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func ConvertMovieDomainToResponse(movieDomain model.MovieDomainInterface) *response.MovieResponse {
	return &response.MovieResponse{
		ID:          movieDomain.GetID(),
		Name:        movieDomain.GetName(),
		Genre:       movieDomain.GetGenre(),
		Description: movieDomain.GetDescription(),
		UrlImg:      movieDomain.GetUrlImg(),
	}
}
