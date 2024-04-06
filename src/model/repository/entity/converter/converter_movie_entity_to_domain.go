package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConvertMovieEntityToDomain(movieEntity entity.MovieEntity) *model.MovieDomainInterface {
	domain := model.NewMovieDomain(movieEntity.Name, movieEntity.Genre, movieEntity.Description, movieEntity.UrlImg)
	domain.SetID(movieEntity.ID)
	return &domain
}
