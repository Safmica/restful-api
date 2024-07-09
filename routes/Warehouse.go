package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func WarehouseRoutes(app *fiber.App) {
	app.Get("/warehouses", controller.GetAllWarehouse)
	app.Get("/warehouses/:id", controller.GetWarehouseByID)
	app.Get("/warehouses/product/:product_id", controller.GetWarehouseByProductID)
	app.Post("/warehouses", controller.CreateWarehouse)
	app.Put("/warehouses/:id", controller.UpdateWarehouse)
}
