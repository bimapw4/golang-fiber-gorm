package main

import (
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Connect()
	config.Migrate(&model.User{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!")
	})

	app.Listen(":3000")
}
