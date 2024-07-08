package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func GetWarehouseByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	warehouseID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || warehouseID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid warehouse_id",
		})
	}

	var warehouse entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").First(&warehouse, warehouseID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Warehouse not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch warehouse",
		})
	}

	return ctx.JSON(fiber.Map{
		"warehouse": warehouse,
	})
}

func GetWarehouseByProductID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("product_id")

	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || productID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product_id",
		})
	}

	var warehouses []entity.Warehouse
	result := database.DB.Preload("Stocks").Preload("Stocks.Product").Preload("Stocks.Product.Category").Joins("JOIN product_warehouses ON warehouses.id = product_warehouses.warehouse_id").Where("product_id = ?", productID).Find(&warehouses)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Warehouse not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch warehouse",
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
