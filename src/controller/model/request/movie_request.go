package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type MovieRequest struct {
	Name        string `json:"name" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	Description string `json:"description" validate:"required"`
	UrlImg      string `json:"url_img" validate:"required"`
}

func (e *MovieRequest) ValidateUserRequest() ([]string, error) {
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
