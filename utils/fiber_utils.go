package utils

import (
    "context"
    "BE_PROJECTUAS/apps/models"

    "github.com/gofiber/fiber/v2"
)

func WrapLogic[Req any, Resp any](
    svcFunc func(context.Context, Req) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        req, parseErr := ParseBody[Req](c)
        if parseErr != nil {
            return parseErr
        }

        ctx := c.Context()
        resp, err := svcFunc(ctx, req)
        if err != nil {
            code := fiber.StatusBadRequest
            if err.Error() == "invalid credentials" {
                code = fiber.StatusUnauthorized
            }
            return ErrorResponse(c, code, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}


// WrapWithUser: For Create (body + userID)
func WrapWithUser[Req any, Resp any](
    svcFunc func(context.Context, Req, string) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID, ok := c.Locals("userID").(string)
        if !ok {
            return ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated")
        }

        req, parseErr := ParseBody[Req](c)
        if parseErr != nil {
            return parseErr
        }

        ctx := c.Context()
        resp, err := svcFunc(ctx, req, userID)
        if err != nil {
            return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}

// WrapParam: For /:id no body (submit, verify, delete)
func WrapParam(
    svcFunc func(context.Context, string) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        if err := svcFunc(c.Context(), id); err != nil {
            return ErrorResponse(c, fiber.StatusConflict, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"status": "updated", "updated_at": GetCurrentTime()})
    }
}

// WrapReject: For reject (body + id, note)
func WrapReject(
    svcFunc func(context.Context, string, string) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        req, parseErr := ParseBody[models.RejectRequest](c)
        if parseErr != nil {
            return parseErr
        }

        if err := svcFunc(c.Context(), id, req.Note); err != nil {
            return ErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"status": "rejected"})
    }
}

// WrapUpdate: For update (body + id) - New, fix mismatch
func WrapUpdate[Req any](
    svcFunc func(context.Context, string, Req) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        req, parseErr := ParseBody[Req](c)
        if parseErr != nil {
            return parseErr
        }

        if err := svcFunc(c.Context(), id, req); err != nil {
            return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"status": "updated"})
    }
}


// WrapListOwn: For list (extract userID, no body)
func WrapListOwn(
    svcFunc func(context.Context, string) ([]models.Achievement, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID, ok := c.Locals("userID").(string)
        if !ok {
            return ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated")
        }

        list, err := svcFunc(c.Context(), userID)
        if err != nil {
            return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
        }
        return SuccessResponse(c, list)
    }
}
// WrapRefresh: For refresh (body with refreshToken)
func WrapRefresh(
    svcFunc func(context.Context, string) (*models.LoginResponse, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        type RefreshReq struct {
            RefreshToken string `json:"refreshToken"`
        }
        req, parseErr := ParseBody[RefreshReq](c)
        if parseErr != nil {
            return parseErr
        }

        resp, err := svcFunc(c.Context(), req.RefreshToken)
        if err != nil {
            return ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}

// WrapProfile: For profile (no body, userID from locals)
func WrapProfile(
    svcFunc func(context.Context, string) (*models.UserResponse, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID, ok := c.Locals("userID").(string)
        if !ok {
            return ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated")
        }

        resp, err := svcFunc(c.Context(), userID)
        if err != nil {
            return ErrorResponse(c, fiber.StatusNotFound, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}

// WrapLogout: Simple no body
func WrapLogout(
    svcFunc func(context.Context) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if err := svcFunc(c.Context()); err != nil {
            return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"message": "Logged out successfully"})
    }
}