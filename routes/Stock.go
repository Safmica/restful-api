package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app *fiber.App) {
	app.Get("/stocks", controller.GetAllStock)
	app.Get("/stocks/:id", controller.GetStockByID)
	app.Get("/stocks/product/:product_id", controller.GetStockByProductID)
	app.Get("/stocks/warehouse/:warehouse_id", controller.GetStockByWarehouseID)
	app.Get("/stocks/warehouse/:warehouse_id/product/:product_id", controller.GetStockByWarehouseIDProductID)
	app.Post("/stocks", controller.CreateStock)
	app.Put("/stocks/:id", controller.UpdateStock)
	app.Put("/stocks/warehouse/:warehouse_id/product/:product_id", controller.UpdateStockByWarehouseIDProductID)
}
