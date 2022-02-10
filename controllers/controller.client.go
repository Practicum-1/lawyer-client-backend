package controllers

import "github.com/gofiber/fiber/v2"

func GetClient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Client got"})
}
