package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Get("/category", controller.GetAllCategory)
	app.Post("/category", controller.CreateCategory)
}
