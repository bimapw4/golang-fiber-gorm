package main

import (
	"golang-fiber-gorm/app/model"
	"golang-fiber-gorm/config"
	"golang-fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Connect()
	config.Migrate(&model.User{})
	config.RedisClient()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!")
	})

	routes.UserRouter(app)
	app.Listen(":3000")
}
