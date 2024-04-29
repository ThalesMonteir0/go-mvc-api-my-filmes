package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
	"time"
)

func (lr *listRepository) CreateList(listDomain model.ListDomainInterface) *rest_err.RestErr {
	var id int
	listEntity := converter.ConvertListDomainToEntity(listDomain)

	err := lr.DB.QueryRow(createListSQL, listEntity.Title, listEntity.Description, time.Now(), listEntity.UserID).Scan(&id)
	if err != nil {
		return rest_err.NewInternalServerError("Unable to create list")
	}

	return nil
}
