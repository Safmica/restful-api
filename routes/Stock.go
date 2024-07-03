package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app *fiber.App) {
	app.Get("/stock", controller.GetAllStock)
	app.Get("/stock/:id", controller.GetStockByID)
	app.Post("/stock", controller.CreateStock)
}
