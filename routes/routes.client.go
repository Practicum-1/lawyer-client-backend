package routes

import (
	"github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(app fiber.Router) error {
	app.Get("/", controllers.GetAllClients)
	app.Get("/:id", controllers.GetClientById)
	app.Post("/", controllers.CreateClient)
	return nil
}
