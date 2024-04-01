package model

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetEmail() string
	SetID(id int)
	GetID() int
	EncryptPassword()
	NameToUpperCase()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		0, email, password, name,
	}
}

func NewUserUpdate(email, name string) UserDomainInterface {
	return &userDomain{
		email: email,
		name:  name,
	}
}

func NewUserLogin(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
