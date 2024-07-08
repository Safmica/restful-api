package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app *fiber.App) {
	app.Get("/stocks", controller.GetAllStock)
	app.Get("/stocks/:id", controller.GetStockByID)
	app.Post("/stocks", controller.CreateStock)
}
