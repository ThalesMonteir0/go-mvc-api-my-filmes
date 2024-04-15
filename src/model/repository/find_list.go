package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (lr *listRepository) FindListsPerUserID(id int) (*[]model.ListDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
