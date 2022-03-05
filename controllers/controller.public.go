package controllers

import (
	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetDashboardData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Client got"})
}

func GetSeededData(c *fiber.Ctx) error {
	type JudiciaryData struct {
		Courts        []models.Court        `json:"courts"`
		Languages     []models.Language     `json:"languages"`
		PracticeAreas []models.PracticeArea `json:"practice_areas"`
		Locations     []models.Location     `json:"location"`
	}
	courts, err := repositories.GetAllCourts()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	languages, err := repositories.GetAllLanguages()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	practiceAreas, err := repositories.GetAllPracticeAreas()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	location, err := repositories.GetAllLocations()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return helpers.SendResponse(c, fiber.StatusOK, "Judiciary Data Fetched", JudiciaryData{Courts: courts, Languages: languages, PracticeAreas: practiceAreas, Locations: location})
}
