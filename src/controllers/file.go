package controllers

import (
	"os"
	"strings"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
)

func DownloadFile(c *fiber.Ctx) error {
	filename := c.Query("filename")
	if filename == "" {
		return utils.Response(c, 400, "[Bad Request]", nil)
	}

	if strings.Contains(filename, "..") || strings.Contains(filename, "\\") || strings.Contains(filename, "/") {
		return utils.Response(c, 400, "[Bad Request]", nil)
	}

	file, err := os.OpenFile("./uploads/"+filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			return utils.Response(c, 404, "[Not Found]", err)
		}
		return utils.Response(c, 500, "[Internal Server Error]", err)
	}
	file.Close()

	return utils.ResponseFile(c, 200, "[Success]", "./uploads/"+filename, filename)
}
