package middleware

import (
    "BE_PROJECTUAS/apps/models"
    "github.com/gofiber/fiber/v2"
)

func RequirePermission(requiredPermission string) fiber.Handler {
    return func(c *fiber.Ctx) error {

        claimsRaw := c.Locals("claims")
        if claimsRaw == nil {
            return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
        }

        claims := claimsRaw.(*models.JwtCustomClaims)

        for _, perm := range claims.Permissions {
            if perm == requiredPermission {
                return c.Next()
            }
        }

        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }
}
