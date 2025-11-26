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
func RequireAnyPermission(perms ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {

        claims := c.Locals("claims").(*models.JwtCustomClaims)
        userPerms := claims.Permissions

        // check if user has ANY of the required permissions
        for _, p := range perms {
            for _, up := range userPerms {
                if p == up {
                    return c.Next()
                }
            }
        }

        return c.Status(403).JSON(fiber.Map{
            "error": "forbidden",
        })
    }
}
