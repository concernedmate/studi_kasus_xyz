package routes

import (
	"studi_kasus_xyz/controllers"
	"studi_kasus_xyz/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/v1/auth", middlewares.AuthLimiter(), controllers.Auth)
}
