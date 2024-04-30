package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConvertListEntityToDomain(listEntity entity.ListEntity) *model.ListDomainInterface {
	listDomain := model.NewListDomain(listEntity.Title, listEntity.Description, listEntity.UserID)
	listDomain.SetID(listEntity.ID)
	return &listDomain
}
