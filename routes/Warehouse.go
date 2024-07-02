package routes

import (
	"GDSC-PROJECT/controller"

	"github.com/gofiber/fiber/v2"
)

func WarehouseRoutes(app *fiber.App) {
	app.Get("/warehouse", controller.GetAllWarehouse)
	app.Post("/warehouse", controller.CreateWarehouse)
}
