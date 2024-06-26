package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func (ud *userDomainService) FindUserByID(id int) (model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := ud.repository.FindUserByID(id)
	if err != nil {
		//	todo:logs
		return nil, err
	}

	return *userDomain, nil
}

func (ud *userDomainService) findUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := ud.repository.FindUserByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}
	return *userDomain, nil
}
