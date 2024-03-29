package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	Name            string `json:"name" validate:"required,min=2"`
}

func (e *UserRequest) ValidateUserRequest() ([]string, error) {
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

type UserUpdateRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required, min=2"`
}
