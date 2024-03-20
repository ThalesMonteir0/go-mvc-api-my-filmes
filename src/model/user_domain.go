package model

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

type UserDomain struct {
	Email           string
	Password        string
	ConfirmPassword string
	Name            string
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUserByID(string) *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
}

func NewUserDomain(email, password, confirmPassword, name string) UserDomainInterface {
	return &UserDomain{
		email, password, confirmPassword, name,
	}
}
