package repository

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

type listRepository struct {
	DB *sql.DB
}

type ListRepository interface {
	FindListsPerUserID(int) (*[]model.ListDomainInterface, *rest_err.RestErr)
	FindListByID(int) (*model.ListDomainInterface, *rest_err.RestErr)
	CreateList(model.ListDomainInterface) *rest_err.RestErr
	DeleteList(int) *rest_err.RestErr
}

func NewListRepository(db *sql.DB) ListRepository {
	return &listRepository{
		db,
	}
}
