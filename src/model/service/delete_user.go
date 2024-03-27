package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUser(id int) *rest_err.RestErr {
	err := ud.repository.DeleteUserRepository(id)
	if err != nil {
		//	todo:logs
		return err
	}

	return nil
}
