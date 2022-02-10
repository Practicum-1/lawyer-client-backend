package routes

import (
	"github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func LawyerRoutes(app fiber.Router) error {
	app.Get("/", controllers.GetLawyer)

	return nil
}
