package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app *fiber.App) {
	app.Get("/stock", controller.GetAllStock)
	app.Post("/stock", controller.CreateStock)
}
