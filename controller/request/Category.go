package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"
	"log"

	"github.com/gofiber/fiber/v2"
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
