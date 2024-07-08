package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(ctx *fiber.Ctx) error {
	var products []entity.Product

	result := database.DB.Preload("Warehouses").Preload("Category").Find(&products)
	if err := validation.QueryResultValidation(result, "product"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"products": products,
	})
}

func GetProductByID(ctx *fiber.Ctx) error {
	productID, err := validation.ParseAndIDValidation(ctx, "id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var product entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").First(&product, productID)
	if err = validation.EntityByIDValidation(result, "product"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"product": product,
	})
}

func GetProductByCategoryID(ctx *fiber.Ctx) error {
	categoryID, err := validation.ParseAndIDValidation(ctx, "category_id", "category")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var products []entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").Where("category_id = ?", categoryID).Find(&products)
	if err = validation.EntityByIDValidation(result, "product"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"products": products,
	})
}

func GetProductByWarehouseID(ctx *fiber.Ctx) error {
	warehouseID, err := validation.ParseAndIDValidation(ctx, "warehouse_id", "warehouse")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var products []entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").Joins("JOIN product_warehouses ON products.id = product_warehouses.product_id").Where("warehouse_id = ?", warehouseID).Find(&products)
	if err = validation.EntityByIDValidation(result, "product"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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

	if err := validation.ProductValidation(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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
