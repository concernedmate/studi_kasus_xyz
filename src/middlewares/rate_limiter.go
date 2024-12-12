package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        1000,
		Expiration: time.Minute * 15,
	})
}

func AuthLimiter() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:                    10,
		Expiration:             time.Minute * 30,
		SkipSuccessfulRequests: true,
	})
}
