package routes

import (
	controller "GDSC-PROJECT/controller/request"
	"GDSC-PROJECT/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Get("/products", controller.GetAllProduct)
	app.Get("/products/:id", controller.GetProductByID)
	app.Get("/products/category/:category_id", controller.GetProductByCategoryID)
	app.Get("/products/warehouse/:warehouse_id", controller.GetProductByWarehouseID)
	app.Post("/products", middleware.Auth, controller.CreateProduct)
	app.Put("/products/:id", middleware.Auth, controller.UpdateProduct)
	app.Delete("/products/:id", middleware.Auth, controller.DeleteProduct)
}
