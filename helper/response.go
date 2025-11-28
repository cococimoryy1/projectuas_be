package helper

import (
    "BE_PROJECTUAS/apps/models" // Fix: Hapus /apps, konsisten dengan project
    "time"

    "github.com/gofiber/fiber/v2"
)

func ParseBody[T any]() fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req T
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{
                "error": "invalid input",
            })
        }

        c.Locals("parsed_body", req)
        return c.Next()
    }
}

func Success(data any) map[string]any {
    return map[string]any{
        "status": "success",
        "data":   data,
    }
}

func Error(code int, message string) map[string]any {
    return map[string]any{
        "status":  "error",
        "message": message,
        "code":    code,
    }
}

// GetUserFromCtx: Extract user dari locals (set oleh AuthRequired middleware)
func CastUser(v any) (*models.User, bool) {
    u, ok := v.(*models.User)
    return u, ok
}

func GetCurrentTime() time.Time {
    return time.Now()
}