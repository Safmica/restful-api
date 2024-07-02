package controller

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllStock(ctx *fiber.Ctx) error {
	var stocks []entity.Stock

	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Find(&stocks)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stocks",
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}

func CreateStock(ctx *fiber.Ctx) error {
	stock := new(entity.Stock)

	if err := ctx.BodyParser(stock); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	database.DB.Debug().Create(&stock)

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"stock":   stock,
	})
}
