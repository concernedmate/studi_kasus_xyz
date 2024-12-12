package controllers

import (
	"studi_kasus_xyz/entities"
	"studi_kasus_xyz/models"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
)

func GetCustDataByID(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	if id < 1 {
		return utils.Response(c, 400, "[Bad request]", "invalid id")
	}

	response, err := models.GetCustDataByID(c.Context(), id)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}
	return utils.Response(c, 200, "[Success]", response)
}

func UpdateCustData(c *fiber.Ctx) error {
	var input entities.CustomerUpdate

	err := c.BodyParser(&input)
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return utils.Response(c, 400, "[Bad Request] Unprocessable", err.(*fiber.Error).Code)
		}
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = utils.ValidateStruct(input)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = models.UpdateCustData(c.Context(), input)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}
	return utils.Response(c, 200, "[Success]", nil)
}
