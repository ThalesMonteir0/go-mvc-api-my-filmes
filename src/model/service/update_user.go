package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ud *userDomainService) UpdateUser(id int, userDomain model.UserDomainInterface) *rest_err.RestErr {
	if user, _ := ud.repository.FindUserByID(id); user == nil {
		return rest_err.NewNotFoundError("User not found!")
	}

	userDomain.NameToUpperCase()

	return ud.repository.UpdateUser(id, userDomain)
}
