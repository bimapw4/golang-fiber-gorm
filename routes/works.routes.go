package routes

import (
	"golang-fiber-gorm/app/controller"

	"github.com/gofiber/fiber/v2"
)

func WorksRoutes(app fiber.Router) {
	r := app.Group("/works")

	r.Get("/", controller.GetWorks)
	r.Post("/", controller.CreateWorks)
	r.Put("/:id", controller.UpdateWorks)
	r.Delete("/:id", controller.DeleteWorksById)
}
