package main

import (
    "log"
    "os"

    "BE_PROJECTUAS/database"
    "BE_PROJECTUAS/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env tidak ditemukan, pakai environment OS")
    }

    // Connect databases
    database.ConnectPostgres()
    database.ConnectMongo()

    // Setup fiber app
    app := fiber.New()

    // ===== ROUTES ===== (Clean, init repo/service di dalam)
    routes.SetupRoutes(app)

    // ===== HEALTH CHECK =====
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":   "ok",
            "postgres": "connected",
            "mongo":    "connected",
        })
    })

    // Read port from env
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "3000"
    }

    log.Println("🚀 Server running on port", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatal(err)
    }
}