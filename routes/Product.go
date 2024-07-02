package routes

import (
	"GDSC-PROJECT/controller"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Get("/product", controller.GetAllProduct)
	app.Post("/product", controller.CreateProduct)
}
