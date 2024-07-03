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

func GetAllCategory(ctx *fiber.Ctx) error {
	var categories []entity.Category

	result := database.DB.Preload("Products").Find(&categories)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch categories",
		})
	}

	return ctx.JSON(fiber.Map{
		"categories": categories,
	})
}

func GetCategoryByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	categoryID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || categoryID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category ID",
		})
	}

	var category entity.Category
	result := database.DB.Preload("Products").First(&category, categoryID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "category not found",
		})
	}

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch category",
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
