package main

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Databaseinit()
	database.DBMigration()
	app := fiber.New()
	routes.WarehouseRoutes(app)
	routes.ProductRoutes(app)
	routes.CategoryRoutes(app)
	routes.StockRoutes(app)
	app.Listen(":8080")
}
