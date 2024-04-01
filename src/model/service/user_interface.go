package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository"
)

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{
		repository: repository,
	}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (int, *rest_err.RestErr)
	UpdateUser(int, model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(int) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(int) *rest_err.RestErr
	LoginUser(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
