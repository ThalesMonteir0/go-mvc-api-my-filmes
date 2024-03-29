package view

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/response"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
)

func ConvertUserDomainToResponse(domain model.UserDomainInterface) *response.UserResponse {
	return &response.UserResponse{
		Id:    domain.GetID(),
		Name:  domain.GetName(),
		Email: domain.GetEmail(),
	}
}
