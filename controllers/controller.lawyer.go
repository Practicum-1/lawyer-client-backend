package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
)

func GetLawyer(c *fiber.Ctx) error {
	return helpers.SendResponse(c, fiber.StatusOK, "Lawyer found", nil)
}
