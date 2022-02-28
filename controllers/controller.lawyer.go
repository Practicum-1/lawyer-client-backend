package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
)

func GetAllLawyer(c *fiber.Ctx) error {

	lawyer, err := repositories.GetAllLawyers()
	if err != nil {
		if err.Error() == "404" {
			return helpers.SendResponse(c, fiber.StatusNotFound, "No lawyer found", nil)
		} else {
			return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), err)
		}
	}
	return helpers.SendResponse(c, fiber.StatusOK, "success", lawyer)
}

func GetLawyerById(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Invalid Id", err)
	}

	var lawyer models.Lawyer
	err = repositories.GetLawyerById(id, &lawyer)
	if err != nil {
		if err.Error() == "404" {
			return helpers.SendResponse(c, fiber.StatusNotFound, "Lawyer not found", nil)
		} else {
			return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), err)
		}
	}
	return helpers.SendResponse(c, fiber.StatusOK, "success", lawyer)

}

func CreateLawyer(c *fiber.Ctx) error {
	newClient := &models.Lawyer{}

	//Parse the body
	if err := c.BodyParser(newClient); err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Invalid input", err)
	}

	//Create Client
	err := repositories.CreateLawyer(newClient)
	fmt.Println("Error: ", err)
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	//Generate token
	fmt.Println("newClient: ", newClient)
	token, err := helpers.GenerateToken(newClient.ID, newClient.Email)

	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Failed to generate token", err)
	}

	//create map string interface

	response := map[string]interface{}{"token": token}

	return helpers.SendResponse(c, fiber.StatusCreated, "Lawyer Created Successfully", response)
}
