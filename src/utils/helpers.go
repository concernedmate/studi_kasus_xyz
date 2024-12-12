package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			return errors.New(" " + err.Field() + " error, " + err.Tag())
		}
	}
	return nil
}

func Response(c *fiber.Ctx, status int, message string, data any) error {
	Logger(c, message)

	if data == nil {
		return c.Status(status).JSON(fiber.Map{
			"message": message,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func ResponseFile(c *fiber.Ctx, status int, message string, filepath string, filename string) error {
	Logger(c, message+" sent file "+filepath)

	split := strings.Split(filename, ".")
	ext := split[len(split)-1]
	if ext == "png" || ext == "jpg" || ext == "jpeg" || ext == "gif" {
		c.Set("Content-Type", "image/"+ext)
	}
	c.Set("Content-Disposition", "attachment; filename="+filename)
	return c.Status(status).SendFile(filepath)
}
