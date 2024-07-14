package routes

import (
	controller "GDSC-PROJECT/controller/request"
	"GDSC-PROJECT/middleware"

	"github.com/gofiber/fiber/v2"
)

func StockRoutes(app *fiber.App) {
	app.Get("/stocks", controller.GetAllStock)
	app.Get("/stocks/:id", controller.GetStockByID)
	app.Get("/stocks/product/:product_id", controller.GetStockByProductID)
	app.Get("/stocks/warehouse/:warehouse_id", controller.GetStockByWarehouseID)
	app.Get("/stocks/warehouse/:warehouse_id/product/:product_id", controller.GetStockByWarehouseIDProductID)
	app.Post("/stocks", middleware.Auth, controller.CreateStock)
	app.Put("/stocks/:id", middleware.Auth, controller.UpdateStock)
	app.Put("/stocks/warehouse/:warehouse_id/product/:product_id", middleware.Auth, controller.UpdateStockByWarehouseIDProductID)
	app.Delete("/stocks/:id", middleware.Auth, controller.DeleteStock)
	app.Delete("/stocks/warehouse/:warehouse_id/product/:product_id", middleware.Auth, controller.DeleteStockByWarehouseIDProductID)
}
