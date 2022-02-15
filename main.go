package main

import (
	"log"
	"os"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	routes.ClientRoutes(app.Group("/client"))
	routes.LawyerRoutes(app.Group("/lawyer"))
	routes.RequestRoutes(app.Group("/request"))
	routes.PublicRoutes(app.Group("/public"))
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	models.Migrate()

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "server running on"})
	})

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
