package middleware

import (
	jwtToken "absent/pkg/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		log.Println("unauthorized", http.StatusBadRequest)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"code":    http.StatusUnauthorized,
			"data":    nil,
			"message": "unauthorized",
		})
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"code":    http.StatusUnauthorized,
			"data":    nil,
			"message": "unauthorized",
		})
	}
	token = parts[1]

	claims, err := jwtToken.DecodeToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"code":    http.StatusUnauthorized,
			"data":    nil,
			"message": "unauthorized",
		})
	}

	c.Locals("userLogin", claims)
	return c.Next()
}
