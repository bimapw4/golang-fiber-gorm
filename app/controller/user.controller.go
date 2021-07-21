package controller

import (
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/app/types"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	user := new(types.UserDTO)

	if err := c.BodyParser(user); err != nil {
		c.Status(403).JSON(err)
		return c.JSON(c)
	}

	model.CreateUser(&model.User{
		Username: user.Username,
		Address:  user.Address,
		Email:    user.Email,
	})
	return c.JSON(user)
}

func UpdateEmailUser(c *fiber.Ctx) error {
	user := new(types.UserUpdateDTO)
	if err := c.BodyParser(user); err != nil {
		c.Status(403).JSON(err)
		return c.JSON(c)
	}

	id := c.Params("id")
	model.UpdateUserById(id, map[string]interface{}{
		"email": user.Email,
	})

	return c.JSON(user)
}
