package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ls *listService) FindListsByUserID(userId int) ([]model.ListDomainInterface, *rest_err.RestErr) {
	listsDomain, err := ls.repository.FindListsPerUserID(userId)
	if err != nil {
		return nil, err
	}

	return *listsDomain, nil
}
