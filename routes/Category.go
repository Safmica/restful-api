package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Get("/categories", controller.GetAllCategory)
	app.Get("/categories/:id", controller.GetCategoryByID)
	app.Get("/categories/product/:product_id", controller.GetCategoryByProductID)
	app.Post("/categories", controller.CreateCategory)
	app.Put("/categories/:id", controller.UpdateCategory)
}
