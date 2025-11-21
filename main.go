package main

import (
	"log"
	"os"

	"BE_PROJECTUAS/database" // GANTI kalau module kamu namanya beda

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env (kalau ada)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env tidak ditemukan, pakai environment OS")
	}

	// Connect ke PostgreSQL
	database.ConnectPostgres()

	// Connect ke MongoDB
	database.ConnectMongo()

	// Setup Fiber
	app := fiber.New()

	// Route test /health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":   "ok",
			"postgres": "connected",
			"mongo":    "connected",
		})
	})

	// Baca port dari env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("🚀 Server running on port", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
