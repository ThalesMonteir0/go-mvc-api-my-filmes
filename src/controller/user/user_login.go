package user

import (
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/controller/model/request"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/model"
	"github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/view"
	"github.com/gofiber/fiber/v2"
)

func (uc *userControllerService) LoginUser(c *fiber.Ctx) error {
	var user request.UserLogin
	errParse := c.BodyParser(&user)
	if errParse != nil {

	}

	userDomain := model.NewUserLogin(user.Email, user.Password)
	userService, token, err := uc.service.LoginUser(userDomain)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}

	c.GetRespHeader("Authorization", token)
	return c.Status(fiber.StatusOK).JSON(view.ConvertUserDomainToResponse(userService))
}
