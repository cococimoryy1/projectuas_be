package wrappers

import (
    "BE_PROJECTUAS/apps/models" // Fix: Hapus /apps, konsisten dengan project
    "time"

    "github.com/gofiber/fiber/v2"
)

func ParseBody[T any](c *fiber.Ctx) (T, error) {
    var req T
    if err := c.BodyParser(&req); err != nil {
        ErrorResponse(c, fiber.StatusBadRequest, "Invalid input data")
        return req, err
    }
    return req, nil
}


func SuccessResponse(c *fiber.Ctx, data interface{}) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": "success",
        "data":   data,
    })
}

func ErrorResponse(c *fiber.Ctx, code int, message string) error {
    return c.Status(code).JSON(fiber.Map{
        "status":  "error",
        "message": message,
        "code":    code,
    })
}

// GetUserFromCtx: Extract user dari locals (set oleh AuthRequired middleware)
func GetUserFromCtx(c *fiber.Ctx) (*models.User, bool) {
    userValue := c.Locals("user") // Fix: Langsung c.Locals(), no locals.Get() atau import middleware/locals
    user, ok := userValue.(*models.User)
    return user, ok
}

func GetCurrentTime() time.Time {
    return time.Now()
}