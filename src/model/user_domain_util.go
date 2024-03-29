package model

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (ud *userDomain) EncryptPassword() {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(ud.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		//	todo:logs
	}
	ud.password = string(hashPassword)
}

func (ud *userDomain) NameToUpperCase() {
	ud.name = strings.ToUpper(ud.name)
}
