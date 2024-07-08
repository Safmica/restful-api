package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Get("/products", controller.GetAllProduct)
	app.Get("/products/:id", controller.GetProductByID)
	app.Get("/products/category/:category_id", controller.GetProductByCategoryID)
	app.Get("/products/warehouse/:warehouse_id", controller.GetProductByWarehouseID)
	app.Post("/products", controller.CreateProduct)
}
