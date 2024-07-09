package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(ctx *fiber.Ctx) error {
	var categories []entity.Category

	result := database.DB.Preload("Products").Find(&categories)
	if err := validation.QueryResultValidation(result, "category"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"categories": categories,
	})
}

func GetCategoryByID(ctx *fiber.Ctx) error {
	categoryID, err := validation.ParseAndIDValidation(ctx, "id", "category")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var category entity.Category
	result := database.DB.Preload("Products").First(&category, categoryID)
	if err = validation.EntityByIDValidation(result, "category"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"category": category,
	})
}

func GetCategoryByProductID(ctx *fiber.Ctx) error {
	productID, err := validation.ParseAndIDValidation(ctx, "product_id", "product")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var category entity.CategoryResponse
	result := database.DB.Debug().Joins("JOIN products ON categories.id = products.category_id").Where("products.id = ?", productID).First(&category)
	if err = validation.EntityByIDValidation(result, "category"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"category": category,
	})
}

func CreateCategory(ctx *fiber.Ctx) error {
	category := new(entity.Category)

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.CategoryValidation(category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Debug().Create(&category)

	return ctx.JSON(fiber.Map{
		"message":  "created successfully",
		"category": category,
	})
}

func UpdateCategory(ctx *fiber.Ctx) error {
	requestCategory := new(entity.CategoryResponse)
	category := new(entity.Category)

	categoryID, err := validation.ParseAndIDValidation(ctx, "id", "category")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&category, categoryID)
	if err = validation.EntityByIDValidation(result, "category"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ctx.BodyParser(requestCategory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.CategoryUpdateValidation(requestCategory, category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Save(&category).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message":  "updated successfully",
		"category": category,
	})
}
