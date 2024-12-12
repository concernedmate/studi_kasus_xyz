package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx, message string) error {
	file, err := os.OpenFile("./logs/"+(time.Now().Local()).Format(time.DateOnly+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("failed to log")
	}
	defer file.Close()

	time := (time.Now().Local()).Format("15:04:05")
	if message != "" {
		message = " => Resp " + message
	}

	_, err = file.WriteString(fmt.Sprintf("%s - [%s]:%s %s %s %s\n", time, c.IP(), c.Port(), c.Method(), c.Path(), message))
	fmt.Printf("%s - [%s]:%s %s %s %s\n", time, c.IP(), c.Port(), c.Method(), c.Path(), message)

	if err != nil {
		return errors.New("failed to log")
	}

	return nil
}

func LoggerMiddleware(c *fiber.Ctx) error {
	Logger(c, "")
	return c.Next()
}
