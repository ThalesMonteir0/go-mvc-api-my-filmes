package repository

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (int, *rest_err.RestErr)
	FindUserByID(id int) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserRepository(id int) *rest_err.RestErr
	UpdateUser(id int, userDomain model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email, password string) (*model.UserDomainInterface, *rest_err.RestErr)
}

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		database: db,
	}

}
