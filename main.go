package main

import (
	"GDSC-PROJECT/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Databseinit()
	app := fiber.New()
	app.Listen(":8080")
}
