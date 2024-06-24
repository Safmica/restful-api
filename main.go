package main

import (
	"GDSC-PROJECT/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Databaseinit()
	database.DBMigration()
	app := fiber.New()
	app.Listen(":8080")
}
