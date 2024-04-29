package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ls *listService) CreateList(listDomain model.ListDomainInterface) *rest_err.RestErr {
	err := ls.repository.CreateList(listDomain)
	if err != nil {
		return err
	}

	return nil
}
