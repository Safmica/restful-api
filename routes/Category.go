package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Get("/categories", controller.GetAllCategory)
	app.Get("/categories/:id", controller.GetCategoryByID)
	app.Post("/categories", controller.CreateCategory)
}
