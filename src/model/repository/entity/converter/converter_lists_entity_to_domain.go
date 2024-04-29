package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConverterListsEntityToDomain(lists []entity.ListEntity) *[]model.ListDomainInterface {
	var listsDomain []model.ListDomainInterface

	for _, listEntity := range lists {
		listDomain := model.NewListDomain(listEntity.Title, listEntity.Description, listEntity.UserID)
		listDomain.SetID(listEntity.ID)

		listsDomain = append(listsDomain, listDomain)
	}

	return &listsDomain
}
