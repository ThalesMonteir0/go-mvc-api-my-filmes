package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConvertListDomainToEntity(listDomain model.ListDomainInterface) *entity.ListEntity {
	return &entity.ListEntity{
		Title:       listDomain.GetTitle(),
		Description: listDomain.GetDescription(),
		UserID:      listDomain.GetUserID(),
	}
}
