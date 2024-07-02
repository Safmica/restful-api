package routes

import (
	"GDSC-PROJECT/controller"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Get("/category", controller.GetAllCategory)
	app.Post("/category", controller.CreateCategory)
}
