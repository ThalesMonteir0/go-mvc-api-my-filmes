package converter

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository/entity"
)

func ConverterUserEntityToDomain(userEntity entity.UserEntity) *model.UserDomainInterface {
	domain := model.NewUserDomain(userEntity.Email, userEntity.Password, userEntity.Name)
	domain.SetID(userEntity.ID)
	return &domain
}
