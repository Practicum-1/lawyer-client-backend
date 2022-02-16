package routes

import (
	"fmt"

	controller "github.com/Practicum-1/lawyer-client-backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(app fiber.Router) error {
	fmt.Printf("%T", app)
	app.Get("/", controller.GetClientById)
	return nil
}
