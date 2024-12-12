package main

import (
	"log"
	"os"
	"studi_kasus_xyz/configs"
	"studi_kasus_xyz/models"
	"studi_kasus_xyz/routes"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}

	err = models.InitDb()
	if err != nil {
		log.Fatal("Failed to init db connection: ", err)
	}

	app := fiber.New(fiber.Config{
		AppName:   "Studi Kasus XYZ",
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Use(configs.CorsMiddleware())
	app.Use(utils.LoggerMiddleware)

	app.Get("/api/version", func(c *fiber.Ctx) error {
		return utils.Response(c, 200, "[Success]", os.Getenv("APP_VER"))
	})

	// routes
	routes.AuthRoutes(app)
	routes.CustomerRoutes(app)

	err = app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
