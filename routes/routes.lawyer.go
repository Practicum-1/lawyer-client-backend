package routes

import (
	"github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func LawyerRoutes(app fiber.Router) error {
	app.Get("/*", controllers.GetAllLawyer)
	app.Get("/:id", controllers.GetLawyerById)
	app.Post("/", controllers.CreateLawyer)

	return nil
}
