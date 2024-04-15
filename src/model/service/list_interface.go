package service

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model/repository"
)

type listService struct {
	repository repository.ListRepository
}

type ListServiceInterface interface {
	FindListsByUserID(int) ([]model.ListDomainInterface, *rest_err.RestErr)
	CreateList(model.ListDomainInterface) *rest_err.RestErr
	DeleteList(int) *rest_err.RestErr
}

func NewListService(repository repository.ListRepository) ListServiceInterface {
	return &listService{
		repository: repository,
	}
}
