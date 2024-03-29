package model

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetEmail() string
	SetID(id int)
	getID() int
	EncryptPassword()
	NameToUpperCase()
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
