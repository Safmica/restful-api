package controller

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllWarehouse(ctx *fiber.Ctx) error {
	var warehouses []entity.Warehouse

	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Find(&warehouses)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch warehouses",
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouses": warehouses,
	})
}

func CreateWarehouse(ctx *fiber.Ctx) error {
	warehouse := new(entity.Warehouse)

	if err := ctx.BodyParser(warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	database.DB.Debug().Create(&warehouse)

	return ctx.JSON(fiber.Map{
		"message":   "created successfully",
		"warehouse": warehouse,
	})
}
