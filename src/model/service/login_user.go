package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ud *userDomainService) LoginUser(domain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	user, err := ud.findUserByEmailAndPassword(domain.GetEmail(), domain.GetPassword())
	if err != nil {

	}

	token, err := user.GenerateToken()
	if err != nil {

	}

	return nil, token, nil
}
