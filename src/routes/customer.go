package routes

import (
	"studi_kasus_xyz/controllers"
	"studi_kasus_xyz/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App) {
	app.Get("/api/v1/customer", middlewares.Limiter(), controllers.GetCustDataByID)
	app.Put("/api/v1/customer", middlewares.Limiter(), controllers.UpdateCustData)

	app.Get("/api/v1/customer/transaction", middlewares.Limiter(), controllers.GetTransactionFromCustID)
	app.Post("/api/v1/customer/transaction", middlewares.Limiter(), controllers.InsertCustTransaction)
}
