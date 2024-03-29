package user

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/gofiber/fiber/v2"
)

func (uc *userControllerService) EditUser(c *fiber.Ctx) error {
	var user request.UserUpdateRequest
	id, _ := c.ParamsInt("id")
	err := c.BodyParser(&user)
	if err != nil {
		rest_err.NewBadRequestError("Fixed some errors with the body")
	}

	userDomain := model.NewUserUpdate(user.Email, user.Name)
	errService := uc.service.UpdateUser(id, userDomain)
	if errService != nil {
		return c.Status(errService.Code).JSON(errService.Message)
	}

	return c.Status(fiber.StatusOK).JSON("ok")
}
