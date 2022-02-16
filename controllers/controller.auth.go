package controllers

import (
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Login "})
}

func SignUp(c *fiber.Ctx) error {
	type NewUser struct {
		Role     string `json:"role"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// user := new(models.Client)
	user := new(NewUser)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	if user.Role == "client" {
		newClient := new(models.Client)
		if err := c.BodyParser(newClient); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
		}
		newClient.Password = hash

		err = repositories.CreateUser(newClient)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create client", "data": err})
		}
		return c.JSON(fiber.Map{"status": "success", "message": "Created Client", "data": newClient})

	} else {
		newLawyer := new(models.Lawyer)
		if err := c.BodyParser(newLawyer); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
		}

		newLawyer.Password = hash

		err = repositories.CreateLayer(newLawyer)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create lawyer", "data": err})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Created lawyer", "data": newLawyer})

	}

}
