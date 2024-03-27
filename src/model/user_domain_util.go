package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(ud.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		//	todo:logs
	}
	ud.password = string(hashPassword)
}
