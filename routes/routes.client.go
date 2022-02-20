package routes

import (
	controller "github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(app fiber.Router) error {
	app.Get("/", controller.GetClientById)
	app.Post("/", controller.CreateClient)
	return nil
}
