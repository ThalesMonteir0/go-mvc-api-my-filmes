package model

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
}

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

func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) EncryptPassword() {
	//	TODO:IMPLEMENTAR

}
func (ud *userDomain) SetID(id string) {
	ud.ID = id
}
func (ud *userDomain) getID() string {
	return ud.ID
}
