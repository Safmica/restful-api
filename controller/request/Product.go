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

func GetProductByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	productID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || productID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product_id",
		})
	}

	var product entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").First(&product, productID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch product",
		})
	}

	return ctx.JSON(fiber.Map{
		"product": product,
	})
}

func GetProductByCategoryID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("category_id")

	categoryID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || categoryID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category_id",
		})
	}

	var products []entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").Where("category_id = ?", categoryID).Find(&products)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch product",
		})
	}

	return ctx.JSON(fiber.Map{
		"products": products,
	})
}

func GetProductByWarehouseID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("warehouse_id")

	warehouseID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || warehouseID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid warehouse_id",
		})
	}

	var products []entity.Product
	result := database.DB.Preload("Warehouses").Preload("Category").Joins("JOIN product_warehouses ON products.id = product_warehouses.product_id").Where("warehouse_id = ?", warehouseID).Find(&products)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch product",
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
