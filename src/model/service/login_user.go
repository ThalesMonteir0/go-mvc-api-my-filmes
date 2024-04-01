package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ud *userDomainService) LoginUser(domain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	return nil, "", nil
}
