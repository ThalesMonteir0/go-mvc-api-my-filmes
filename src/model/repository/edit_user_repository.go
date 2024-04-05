package repository

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity/converter"
	"time"
)

func (ur *userRepository) UpdateUser(userID int, userDomain model.UserDomainInterface) *rest_err.RestErr {
	userEntity := converter.ConverterUserDomainToEntity(userDomain)
	var id int

	if err := ur.database.QueryRow(sqlUpdateUser, userEntity.Name, userEntity.Email, time.Now(), userID).Scan(&id); err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}
	if id == 0 {
		return rest_err.NewInternalServerError("An error occurred while updating the user")
	}

	return nil
}
