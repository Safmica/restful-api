package routes

import (
	controller "GDSC-PROJECT/controller/request"
	"GDSC-PROJECT/middleware"

	"github.com/gofiber/fiber/v2"
)

func WarehouseRoutes(app *fiber.App) {
	app.Get("/warehouses", controller.GetAllWarehouse)
	app.Get("/warehouses/:id", controller.GetWarehouseByID)
	app.Get("/warehouses/product/:product_id", controller.GetWarehouseByProductID)
	app.Post("/warehouses", middleware.Auth, controller.CreateWarehouse)
	app.Put("/warehouses/:id", middleware.Auth, controller.UpdateWarehouse)
	app.Delete("/warehouses/:id", middleware.Auth, controller.DeleteWarehouse)
}
