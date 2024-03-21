package model

type userDomain struct {
	Email           string
	Password        string
	ConfirmPassword string
	Name            string
}

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetConfirmPassword() string
	GetEmail() string
	EncryptPassword()
}

func NewUserDomain(email, password, confirmPassword, name string) UserDomainInterface {
	return &userDomain{
		email, password, confirmPassword, name,
	}
}

func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetConfirmPassword() string {
	return ud.ConfirmPassword
}
func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) EncryptPassword() {
	//	TODO:IMPLEMENTAR

}
