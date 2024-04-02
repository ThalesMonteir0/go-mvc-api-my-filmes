package model

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
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

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(int),
		email: claims["email"].(string),
		name:  claims["name"].(string),
	}, nil
}
func VerifyTokenMiddleware(c *fiber.Ctx) error {
	secret := os.Getenv(JWT_SECRET)
	tokenValue := c.Get("Authorization")

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	})

	if err != nil {
		err := rest_err.NewUnauthorizedRequestError("invalid token")
		return c.Status(err.Code).JSON(err.Message)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err := rest_err.NewUnauthorizedRequestError("invalid token")
		return c.Status(err.Code).JSON(err.Message)
	}

	return c.Next()
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
