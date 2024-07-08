package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllStock(ctx *fiber.Ctx) error {
	var stocks []entity.Stock

	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Find(&stocks)
	if err := validation.QueryResultValidation(result, "stock"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}

func GetStockByID(ctx *fiber.Ctx) error {
	stockID, err := validation.ParseAndIDValidation(ctx, "id", "stock")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var stock entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").First(&stock, stockID)
	if err = validation.EntityByIDValidation(result, "stock"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"stock": stock,
	})
}

func GetStockByProductID(ctx *fiber.Ctx) error {
	productID, err := validation.ParseAndIDValidation(ctx, "product_id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("product_id = ?", productID).Find(&stocks)
	if err = validation.EntityByIDValidation(result, "stock"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}
func GetStockByWarehouseID(ctx *fiber.Ctx) error {
	warehouseID, err := validation.ParseAndIDValidation(ctx, "warehouse_id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("warehouse_id = ?", warehouseID).Find(&stocks)
	if err = validation.EntityByIDValidation(result, "stock"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"stocks": stocks,
	})
}

func GetStockByWarehouseIDProductID(ctx *fiber.Ctx) error {
	warehouseID, err := validation.ParseAndIDValidation(ctx, "warehouse_id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productID, err := validation.ParseAndIDValidation(ctx, "product_id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var stocks []entity.Stock
	result := database.DB.Preload("Product").Preload("Product.Category").Preload("Warehouse").Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).Find(&stocks)
	if err = validation.EntityByIDValidation(result, "stock"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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
