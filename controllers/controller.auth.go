package controllers

import (
	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
	"github.com/gofiber/fiber/v2"
)

func ClientLogin(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var client models.Client

	//Parse the body
	if err := c.BodyParser(&input); err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Invalid input", err)
	}

	// check if user exist
	err := repositories.GetUserByEmail(&client, input.Email)

	if err != nil {
		return helpers.SendResponse(c, fiber.StatusUnauthorized, "Invalid Email or Password", nil)
	}

	// check if password is correct
	if !helpers.CheckPasswordHash(input.Password, client.Password) {
		return helpers.SendResponse(c, fiber.StatusUnauthorized, "Invalid Email or Password", nil)
	}

	//Generate token
	token, err := helpers.GenerateToken(client.ID, client.Email)
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Failed to generate token", err)
	}

	//create map string interface
	response := map[string]interface{}{"token": token, "client": client}
	return helpers.SendResponse(c, fiber.StatusCreated, "Successfully Logged in", response)

}
