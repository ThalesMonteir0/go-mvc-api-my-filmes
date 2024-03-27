package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (int, *rest_err.RestErr) {
	var id int
	userEntity := converter.ConverterUserDomainToEntity(userDomain)
	if err := ur.database.QueryRow(sqlCreateUser, userEntity.Email, userEntity.Password, userEntity.Name).Scan(&id); err != nil {
		//	TODO:LOGS
		return id, rest_err.NewInternalServerError(err.Error())
	}
	return id, nil
}
