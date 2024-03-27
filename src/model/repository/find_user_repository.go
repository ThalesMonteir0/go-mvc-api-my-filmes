package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
)

func (ur *userRepository) FindUserByID(id int) (*model.UserDomainInterface, *rest_err.RestErr) {
	var user entity.UserEntity

	if err := ur.database.QueryRow(sqlFindUserByID, id).Scan(&user); err != nil {
		//	todo:logs
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ConverterUserEntityToDomain(user), nil
}
