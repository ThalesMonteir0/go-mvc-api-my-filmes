package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConverterUserDomainToEntity(userDomain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
	}
}
