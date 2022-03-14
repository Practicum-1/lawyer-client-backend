package routes

import (
	"github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RequestRoutes(app fiber.Router) error {
	app.Get("/:id", controllers.GetRequest)
	app.Post("/", controllers.CreateRequest)
	app.Delete("/:id", controllers.DeleteRequest)
	app.Patch("/:status/:id", controllers.ChangeRequestStatus)
	return nil
}
