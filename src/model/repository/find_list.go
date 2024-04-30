package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (lr *listRepository) FindListsPerUserID(id int) (*[]model.ListDomainInterface, *rest_err.RestErr) {
	var lists []entity.ListEntity

	rows, err := lr.DB.Query(findListsByUserIDSQL, id)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Unable to list movie lists for this user")
	}

	for rows.Next() {
		var list entity.ListEntity
		if err = rows.Scan(&list.ID, &list.Title, &list.Description, &list.Created_at, &list.Deleted_at, &list.Updated_at, &list.UserID); err != nil {
			return nil, rest_err.NewInternalServerError("Unable to list movie lists for this user")
		}
	}

	return converter.ConverterListsEntityToDomain(lists), nil
}

func (lr *listRepository) FindListByID(listID int) (*model.ListDomainInterface, *rest_err.RestErr) {
	var listEntity entity.ListEntity

	if err := lr.DB.QueryRow(findListByIDSQL, listID).
		Scan(&listEntity.ID, listEntity.Title, &listEntity.Description, &listEntity.Created_at,
			&listEntity.Deleted_at, &listEntity.Updated_at, &listEntity.UserID); err != nil {
		return nil, rest_err.NewBadRequestError("list not found")
	}

	return converter.ConvertListEntityToDomain(listEntity), nil
}
