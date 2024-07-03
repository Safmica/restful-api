package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func WarehouseRoutes(app *fiber.App) {
	app.Get("/warehouse", controller.GetAllWarehouse)
	app.Get("/warehouse/:id", controller.GetWarehouseByID)
	app.Post("/warehouse", controller.CreateWarehouse)
}
