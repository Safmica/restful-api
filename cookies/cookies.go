package cookies

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetJwtCookie(ctx *fiber.Ctx, token string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 2),
		HTTPOnly: true,
	})
}

func GetJwtFromCookie(ctx *fiber.Ctx) string {
	return ctx.Cookies("jwt")
}

func ClearJwtCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
}
