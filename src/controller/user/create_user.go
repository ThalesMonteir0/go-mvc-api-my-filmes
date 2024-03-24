package user

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/view"
	"github.com/gofiber/fiber/v2"
)

func (uc *userControllerService) CreateUser(c *fiber.Ctx) error {
	user := request.UserRequest{}
	if err := c.BodyParser(&user); err != nil {
		newErr := rest_err.NewBadRequestError("Invalid parametros")
		return c.Status(newErr.Code).JSON(newErr.Message)
	}

	userDomain := model.NewUserDomain(user.Email, user.Password, user.ConfirmPassword, user.Name)
	uc.service.CreateUser(userDomain)

	//err := user.ValidateUserRequest()
	//if err != nil {
	//	newErr := rest_err.NewBadRequestError("Invalid parametros")
	//	return c.Status(newErr.Code).JSON(newErr.Message)
	//}

	return c.Status(fiber.StatusCreated).JSON(view.ConvertUserDomainToResponse(userDomain))
}
