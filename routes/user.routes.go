package routes

import (
	"golang-fiber-gorm/app/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	r := app.Group("/user")

	r.Post("/", controller.CreateUser)
	r.Put("/:id", controller.UpdateEmailUser)

}
