package model

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	JWT_SECRET = "JWT_SECRET"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"name":  ud.name,
		"email": ud.email,
		"exp":   time.Now().Add(time.Hour * 3),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeString, err := token.SignedString(secret)
	if err != nil {
		return "", rest_err.NewInternalServerError("error trying to generate jwt token")
	}

	return tokeString, nil
}
