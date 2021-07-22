package controller

import (
	"context"
	"encoding/json"
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/app/types"
	"golang-fiber-gorm/config"

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
		Work_id:  user.Work_id,
	})

	config.Client.Del(context.Background(), "GetUser")
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

	config.Client.Del(context.Background(), "GetUser")

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	var user []types.UserListDB
	model.GetUserList(&user)

	dataSave, _ := json.Marshal(types.UserListResp{Status: true, Data: user})

	_ = config.Client.Set(context.Background(), "GetUser", dataSave, 0).Err()

	data, err := config.Client.Get(context.Background(), "GetUser").Result()
	if err != nil {
		return err
	}

	var out map[string]interface{}
	_ = json.Unmarshal([]byte(data), &out)

	return c.JSON(out)
}

func DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	model.DeleteUser("id = " + id)

	config.Client.Del(context.Background(), "GetUser")

	return c.SendStatus(200)
}
