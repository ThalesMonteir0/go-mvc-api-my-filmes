package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (e *UserLogin) ValidateUserRequest() ([]string, error) {
	err := validator.New().Struct(e)
	if err != nil {
		var allErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			allErrors = append(allErrors, fmt.Sprintf("Field %s: %s", err.Field(), err.Tag()))
		}
		return allErrors, nil
	}
	return nil, nil
}
