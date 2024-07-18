package controller

import (
	"GDSC-PROJECT/controller/validation"
	"GDSC-PROJECT/cookies"
	"GDSC-PROJECT/database"
	"GDSC-PROJECT/models/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var users []entity.UserResponse

	result := database.DB.Find(&users)
	if err := validation.QueryResultValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"users": users,
	})
}

func GetUserByID(ctx *fiber.Ctx) error {
	userID, err := validation.ParseAndIDValidation(ctx, "id", "user")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user entity.UserResponse
	result := database.DB.First(&user, userID)
	if err = validation.EntityByIDValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"users": user,
	})
}

func UserLogin(ctx *fiber.Ctx) error {
	user := new(entity.UserLogin)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := validation.UserLoginValidation(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookies.SetJwtCookie(ctx, token)

	return ctx.JSON(fiber.Map{
		"message": "login successful",
	})
}

func UserLogout(ctx *fiber.Ctx) error {
	cookies.ClearJwtCookie(ctx)
	return ctx.JSON(fiber.Map{
		"message": "logout successful",
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(entity.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.UserValidation(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Debug().Create(&user)

	return ctx.JSON(fiber.Map{
		"message": "created successfully",
		"user":    user,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	requestUser := new(entity.UserUpdateResponse)
	user := new(entity.User)

	userID, err := validation.ParseAndIDValidation(ctx, "id", "user")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&user, userID)
	if err = validation.EntityByIDValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ctx.BodyParser(requestUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validation.UserUpdateValidation(requestUser, user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "updated successfully",
		"user":    user,
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := new(entity.User)

	userID, err := validation.ParseAndIDValidation(ctx, "id", "user")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := database.DB.First(&user, userID)
	if err = validation.EntityByIDValidation(result, "user"); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	errUpdate := database.DB.Debug().Delete(&user).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "delete successfully",
		"user":    user,
	})
}
