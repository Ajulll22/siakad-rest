package user

import (
	"siakad-rest/config"
	"siakad-rest/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h handler) GetUsers(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	if claims.Id != "10" {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}
