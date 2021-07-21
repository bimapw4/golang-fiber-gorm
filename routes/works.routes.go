package routes

import (
	"golang-fiber-gorm/app/controller"

	"github.com/gofiber/fiber/v2"
)

func WorksRoutes(app fiber.Router) {
	r := app.Group("/user")

	r.Get("/", controller.GetUser)
	r.Post("/", controller.CreateUser)
	r.Put("/:id", controller.UpdateEmailUser)
	r.Delete("/:id", controller.DeleteUserById)
}
