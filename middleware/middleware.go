package middleware

import (
	"GDSC-PROJECT/cookies"
	"GDSC-PROJECT/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := cookies.GetJwtFromCookie(ctx)
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeJwt(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	exp := claims["exp"].(float64)
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token is expired",
		})
	}

	return ctx.Next()
}
