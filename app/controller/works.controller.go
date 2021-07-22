package controller

import (
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/app/types"

	"github.com/gofiber/fiber/v2"
)

func CreateWorks(c *fiber.Ctx) error {

	work := new(types.WorksDTO)

	if err := c.BodyParser(work); err != nil {
		c.Status(403).JSON(err)
		return c.JSON(c)
	}

	model.CreateUser(&model.Works{
		Work:        work.Work,
		Description: work.Description,
	})
	return c.JSON(work)
}

func UpdateWorks(c *fiber.Ctx) error {
	work := new(types.WorksUpdateDTO)
	if err := c.BodyParser(work); err != nil {
		c.Status(403).JSON(err)
		return c.JSON(c)
	}

	id := c.Params("id")
	model.UpdateWorksById(id, map[string]interface{}{
		"work": work.Work,
	})

	return c.JSON(work)
}

func GetWorks(c *fiber.Ctx) error {
	var works []types.WorksListDB
	model.GetWorksList(&works)

	return c.JSON(types.WorksListResp{
		Status: true,
		Data:   works,
	})
}

func DeleteWorksById(c *fiber.Ctx) error {
	id := c.Params("id")

	model.DeleteWorks("id = " + id)

	return c.SendStatus(200)
}
