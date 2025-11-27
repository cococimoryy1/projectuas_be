package middleware

import (
    "BE_PROJECTUAS/utils"
    "github.com/gofiber/fiber/v2"
    "strings"
)

func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(401).JSON(fiber.Map{"error": "missing token"})
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(401).JSON(fiber.Map{"error": "invalid token format"})
        }

        claims, err := utils.ValidateToken(parts[1])
        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
        }

        // SIMPAN CLAIMS
        c.Locals("claims", claims)
        c.Locals("userID", claims.UserID)
        c.Locals("studentID", claims.StudentID)
        c.Locals("lecturerID", claims.LecturerID)
        c.Locals("role", claims.RoleName)

        // ⬇️ WAJIB DITAMBAHKAN
        c.Locals("permissions", claims.Permissions)

        return c.Next()
    }
}
