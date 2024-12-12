package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
