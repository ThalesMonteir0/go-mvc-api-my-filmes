package model

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetEmail() string
	SetID(id string)
	getID() string
	EncryptPassword()
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		"", email, password, name,
	}
}
