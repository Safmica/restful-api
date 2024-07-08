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

func GetStockByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	stockID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || stockID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid stock ID",
		})
	}

	var stock entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").First(&stock, stockID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "stock not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stock",
		})
	}

	return ctx.JSON(fiber.Map{
		"stock": stock,
	})
}

func GetStockByProductID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("product_id")

	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || productID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product_id",
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("product_id = ?", productID).Find(&stocks)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "stock not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stock",
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}
func GetStockByWarehouseID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("warehouse_id")

	warehouseID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || warehouseID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid warehouse_id",
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("warehouse_id = ?", warehouseID).Find(&stocks)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "stock not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stock",
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}

func GetStockByWarehouseIDProductID(ctx *fiber.Ctx) error {
	idWarehouseParam := ctx.Params("warehouse_id")
	idProductParam := ctx.Params("product_id")

	warehouseID, err := strconv.ParseUint(idWarehouseParam, 10, 64)
	if err != nil || warehouseID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid warehouse_id",
		})
	}

	productID, err := strconv.ParseUint(idProductParam, 10, 64)
	if err != nil || productID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product_id",
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).Find(&stocks)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "stock not found",
			})
		}
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch stock",
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

	if err := validation.StockValidation(stock); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Debug().Create(&stock)

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"stock":   stock,
	})
}
