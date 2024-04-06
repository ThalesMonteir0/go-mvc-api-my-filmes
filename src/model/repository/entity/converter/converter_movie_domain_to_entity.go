package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConvertMovieDomainToEntity(domain model.MovieDomainInterface) *entity.MovieEntity {
	return &entity.MovieEntity{
		Name:        domain.GetName(),
		UrlImg:      domain.GetUrlImg(),
		Genre:       domain.GetGenre(),
		Description: domain.GetDescription(),
	}
}
