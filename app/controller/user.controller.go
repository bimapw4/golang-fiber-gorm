package controller

import (
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/app/types"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	user := new(types.UserDTO)

	if err := c.BodyParser(user); err != nil {
		c.Status(503).JSON(err)
		return c.JSON(c)
	}

	model.CreateUser(&model.User{
		Username: user.Username,
		Address:  user.Address,
		Email:    user.Email,
	})
	return c.JSON(user)
}
