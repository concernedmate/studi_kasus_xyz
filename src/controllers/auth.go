package controllers

import (
	"studi_kasus_xyz/entities"
	"studi_kasus_xyz/models"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	var auth_input entities.Auth

	err := c.BodyParser(&auth_input)
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return utils.Response(c, 400, "[Bad Request] Unprocessable", err.(*fiber.Error).Code)
		}
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = utils.ValidateStruct(auth_input)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	response, err := models.Auth(auth_input)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}
	return utils.Response(c, 200, "[Success]", fiber.Map{
		"username":     response.Username,
		"access_token": response.AccessToken,
		"grup":         response.Grup,
	})
}

func ChangePass(c *fiber.Ctx) error {
	var change_pass_input entities.ChangePass

	err := c.BodyParser(&change_pass_input)
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return utils.Response(c, 400, "[Bad Request] Unprocessable", err.(*fiber.Error).Code)
		}
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = utils.ValidateStruct(change_pass_input)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = models.ChangePass(change_pass_input)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}
	return utils.Response(c, 200, "[Success]", nil)
}
