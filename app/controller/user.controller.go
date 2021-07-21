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

func GetUser(c *fiber.Ctx) error {
	var user []types.UserListDB
	model.GetUserList(&user)

	return c.JSON(types.UserListResp{
		Status: true,
		Data:   user,
	})
}

func DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	model.DeleteUser("id = " + id)

	return c.SendStatus(200)
}
