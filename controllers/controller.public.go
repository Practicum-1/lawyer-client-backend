package controllers

import (
	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetDashboardData(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Client got"})
}

func GetSeededData(c *fiber.Ctx) error {
	response := make(map[string]interface{})

	courts, err := repositories.GetAllCourts()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	//loop on courts
	parseCourts := make(map[uint]interface{})
	for _, court := range courts {
		parseCourts[court.ID] = court.Name
	}
	response["courts"] = parseCourts

	languages, err := repositories.GetAllLanguages()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	//loop on languages
	parseLanguages := make(map[uint]interface{})
	for _, language := range languages {
		parseLanguages[language.ID] = language.Name
	}
	response["languages"] = parseLanguages

	practiceAreas, err := repositories.GetAllPracticeAreas()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	//loop on practice areas
	parsePracticeAreas := make(map[uint]interface{})
	for _, practiceArea := range practiceAreas {
		parsePracticeAreas[practiceArea.ID] = practiceArea.Name
	}
	response["practice_areas"] = parsePracticeAreas
	location, err := repositories.GetAllLocations()
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	//loop on locations
	parseCities := make(map[uint]interface{})
	parseState := make(map[string]interface{})
	for _, loc := range location {
		parseCities[loc.ID] = loc.City
		parseState[loc.State] = loc.State
	}
	response["cities"] = parseCities
	response["states"] = parseState

	return helpers.SendResponse(c, fiber.StatusOK, "Judiciary Data Fetched", response)
}
