package routes

import (
	controller "GDSC-PROJECT/controller/request"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controller.GetAllUser)
	app.Get("/users/:id", controller.GetUserByID)
	app.Post("/users", controller.CreateUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)
}
