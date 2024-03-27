package model

type userDomain struct {
	id       string
	email    string
	password string
	name     string
}

func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) SetID(id string) {
	ud.id = id
}
func (ud *userDomain) getID() string {
	return ud.id
}
