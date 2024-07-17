package routes

import (
	controller "GDSC-PROJECT/controller/request"
	"GDSC-PROJECT/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", middleware.Auth, controller.GetAllUser)
	app.Get("/users/:id", middleware.Auth, controller.GetUserByID)
	app.Post("/users/login", controller.UserLogin)
	app.Post("users/logout", controller.UserLogout)
	app.Post("/users", middleware.Auth, controller.CreateUser)
	app.Put("/users/:id", middleware.Auth, controller.UpdateUser)
	app.Delete("/users/:id", middleware.Auth, controller.DeleteUser)
}
