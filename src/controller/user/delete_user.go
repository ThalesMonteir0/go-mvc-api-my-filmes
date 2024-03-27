package user

import "github.com/gofiber/fiber/v2"

func (uc *userControllerService) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := uc.service.DeleteUser(id)
	if err != nil {
		return c.Status(err.Code).JSON(err.Message)
	}
	return nil
}
