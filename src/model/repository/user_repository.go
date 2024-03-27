package repository

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (ud model.UserDomainInterface, err *rest_err.RestErr)
}

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		database: db,
	}

}
