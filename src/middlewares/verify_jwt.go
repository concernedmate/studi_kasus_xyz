package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"studi_kasus_xyz/configs"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	split := strings.Split(header, " ")

	if len(split) != 2 || header == "" {
		return utils.Response(c, 401, "[Unauthorized]", nil)
	}

	token, err := jwt.Parse(split[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GetAccessKey()), nil
	})
	if err != nil {
		fmt.Println(err)
		return utils.Response(c, 401, "[Unauthorized] Failed to auth", nil)
	}

	if !token.Valid {
		return utils.Response(c, 401, "[Unauthorized] Failed to auth, invalid token!", nil)
	}

	userdata := token.Claims.(jwt.MapClaims)
	username, ok := userdata["username"].(string)
	if !ok {
		return utils.Response(c, 401, "[Unauthorized] Failed to auth, invalid token!", nil)
	}

	grup, ok := userdata["grup"].(string)
	if !ok {
		return utils.Response(c, 401, "[Unauthorized] Failed to auth, invalid token!", nil)
	}

	c.Locals("userdata", map[string]string{
		"username": username,
		"grup":     grup,
	})

	return c.Next()
}

func GetVerifiedUsers(c *fiber.Ctx) (username string, err error) {
	if c.Locals("userdata") == nil {
		return "", errors.New("jwt not found")
	}

	userdata := c.Locals("userdata").(map[string]string)

	if userdata["username"] == "" {
		return "", errors.New("jwt username not found")
	}
	return userdata["username"], nil
}
