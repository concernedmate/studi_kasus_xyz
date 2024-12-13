package routes

import (
	"studi_kasus_xyz/controllers"
	"studi_kasus_xyz/middlewares"

	"github.com/gofiber/fiber/v2"
)

func FileRoutes(app *fiber.App) {
	app.Get("/api/v1/file", middlewares.Limiter(), controllers.DownloadFile)
}
