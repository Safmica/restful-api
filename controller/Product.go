package controller

import (
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(ctx *fiber.Ctx) error {
	var products []entity.Product

	result := database.DB.Preload("Warehouses").Preload("Category").Find(&products)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	return ctx.JSON(fiber.Map{
		"products": products,
	})
}

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(entity.Product)

	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	database.DB.Debug().Create(&product)

	for _, warehouseID := range product.WarehouseID {
		productWarehouse := new(entity.ProductWarehouse)
		productWarehouse.WarehouseID = warehouseID
		productWarehouse.ProductID = product.ID
		database.DB.Create(&productWarehouse)
	}

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"product": product,
	})
}
