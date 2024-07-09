package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllWarehouse(ctx *fiber.Ctx) error {
	var warehouses []entity.Warehouse

	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Find(&warehouses)
	if err := validation.QueryResultValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouses": warehouses,
	})
}

func GetWarehouseByID(ctx *fiber.Ctx) error {
	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var warehouse entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouse": warehouse,
	})
}

func GetWarehouseByProductID(ctx *fiber.Ctx) error {
	productID, err := validation.ParseAndIDValidation(ctx, "product_id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var warehouses []entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Joins("JOIN product_warehouses ON warehouses.id = product_warehouses.warehouse_id").Where("product_id = ?", productID).Find(&warehouses)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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

	if err := validation.WarehouseValidation(warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Debug().Create(&warehouse)

	return ctx.JSON(fiber.Map{
		"message":   "created successfully",
		"warehouse": warehouse,
	})
}

func UpdateWarehouse(ctx *fiber.Ctx) error {
	requestWarehouse := new(entity.WarehouseResponse)
	warehouse := new(entity.Warehouse)

	warehouseID, err := validation.ParseAndIDValidation(ctx, "id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&warehouse, warehouseID)
	if err = validation.EntityByIDValidation(result, "warehouse"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ctx.BodyParser(requestWarehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.WarehouseUpdateValidation(requestWarehouse, warehouse); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Save(&warehouse).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":   "updated successfully",
		"warehouse": warehouse,
	})
}
