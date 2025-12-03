package main

import (
    "log"
    "os"

    "BE_PROJECTUAS/database"
    "BE_PROJECTUAS/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    
    // Swagger
    _ "BE_PROJECTUAS/docs" 
    _ "BE_PROJECTUAS/apps/swagger"
    swagger "github.com/gofiber/swagger"
)

// @title Sistem Pelaporan Prestasi Mahasiswa API
// @version 1.0
// @description Backend REST API untuk pelaporan prestasi, verifikasi beberapa role, dan RBAC lengkap.
// @host localhost:3000
// @BasePath /api/v1

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env tidak ditemukan, pakai environment OS")
    }

    database.ConnectPostgres()
    database.ConnectMongo()

    app := fiber.New()

    // =======================
    // SWAGGER DOC ENDPOINT
    // =======================
    app.Get("/swagger/*", swagger.HandlerDefault) 
    // http://localhost:3000/swagger/index.html

    // =======================
    // ROUTES
    // =======================
    routes.SetupRoutes(app)

    // HEALTH CHECK
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":   "ok",
            "postgres": "connected",
            "mongo":    "connected",
        })
    })

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "3000"
    }

    log.Println(" Server running on port", port)
    app.Listen(":" + port)
}
