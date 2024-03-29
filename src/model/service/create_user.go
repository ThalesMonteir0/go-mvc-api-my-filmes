package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (int, *rest_err.RestErr) {
	if user, _ := ud.repository.FindUserByEmail(userDomain.GetEmail()); user != nil {
		return 0, rest_err.NewBadRequestError("email already exists")
	}

	userDomain.EncryptPassword()
	userDomain.NameToUpperCase()

	id, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		//	TODO: LOGS
		return id, err
	}

	return id, nil
}
