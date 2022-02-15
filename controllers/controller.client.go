package controllers

import (
	"strconv"

	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetClientById(c *fiber.Ctx) error {
	id := c.Params("id")

	intId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "response": "Invalid Id"})
	}

	var client models.Client
	err = repositories.GetClientById(intId, &client)
	if err != nil {
		if err.Error() == "404" {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "msg": "Client not found"})
		} else {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "msg": "Failed to fetch client"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": client})

}

/* a function which accepts 3 parameters
1) fiber Status
2) msg
3) data ---> Struct ---> marshalJSON send
*/
