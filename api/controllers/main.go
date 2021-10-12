package controllers

import (
  "github.com/gofiber/fiber/v2"
)

func GetMain(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message": "Esto es la página principal",
	})
}