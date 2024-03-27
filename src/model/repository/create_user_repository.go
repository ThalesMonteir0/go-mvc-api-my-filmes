package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

var (
	sqlCreateUser = `INSERT INTO users (email, password, name) 
					values ($1,$2,$3); returning id`
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	var id int
	userEntity := converter.ConverterUserDomainToEntity(userDomain)
	if err := ur.database.QueryRow(sqlCreateUser, userEntity.Email, userEntity.Password, userEntity.Name).Scan(&id); err != nil {
		//	logs
	}
	return nil, nil

}
