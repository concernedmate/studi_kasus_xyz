package configs

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() func(*fiber.Ctx) error {
	env := os.Getenv("ENVIRONMENT")
	if env == "PROD" {
		return cors.New(cors.Config{
			AllowOrigins:     "https://iot-trax.com,http://localhost:" + os.Getenv("APP_PORT") + ",https://localhost:" + os.Getenv("APP_PORT"),
			ExposeHeaders:    "Content-Disposition",
			AllowCredentials: true,
		})
	} else {
		return cors.New(cors.Config{
			AllowOrigins:     "http://localhost:3000,http://localhost:4000,https://localhost:3000,https://localhost:4000",
			ExposeHeaders:    "Content-Disposition",
			AllowCredentials: true,
		})
	}
}
