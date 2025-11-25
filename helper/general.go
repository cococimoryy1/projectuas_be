package wrappers

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

        resp, err := svcFunc(c.Context(), req)
        if err != nil {
            return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}

func WrapParam(
    svcFunc func(context.Context, string) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        if err := svcFunc(c.Context(), id); err != nil {
            return ErrorResponse(c, fiber.StatusConflict, err.Error())
        }
        return SuccessResponse(c, fiber.Map{
            "status": "updated",
        })
    }
}

func WrapReject(
    svcFunc func(context.Context, string, string) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")

        req, err := ParseBody[models.RejectRequest](c)
        if err != nil {
            return err
        }

        if err := svcFunc(c.Context(), id, req.Note); err != nil {
            return ErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"status": "rejected"})
    }
}
// WrapRefresh: For refresh (body with refreshToken)

func WrapListAll(
    svcFunc func(context.Context) ([]models.Achievement, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {

        list, err := svcFunc(c.Context())
        if err != nil {
            return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
        }

        return SuccessResponse(c, list)
    }
}

func WrapUpdate[Req any](
    svc func(context.Context, string, Req) error,
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        req, err := ParseBody[Req](c)
        if err != nil {
            return err
        }
        if err := svc(c.Context(), id, req); err != nil {
            return ErrorResponse(c, 400, err.Error())
        }
        return SuccessResponse(c, fiber.Map{"status": "updated"})
    }
}
func WrapLogicParam[Req any, Resp any](
    svc func(context.Context, string, Req) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        req, err := ParseBody[Req](c)
        if err != nil {
            return err
        }
        resp, svcErr := svc(c.Context(), id, req)
        if svcErr != nil {
            return ErrorResponse(c, 400, svcErr.Error())
        }
        return SuccessResponse(c, resp)
    }
}
func WrapParamResp[Resp any](
    svcFunc func(context.Context, string) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")

        resp, err := svcFunc(c.Context(), id)
        if err != nil {
            return ErrorResponse(c, 400, err.Error())
        }

        return SuccessResponse(c, resp)
    }
}
func WrapUpdateResp[Req any, Resp any](
    svcFunc func(context.Context, string, Req) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")

        req, err := ParseBody[Req](c)
        if err != nil {
            return err
        }

        resp, svcErr := svcFunc(c.Context(), id, req)
        if svcErr != nil {
            return ErrorResponse(c, 400, svcErr.Error())
        }

        return SuccessResponse(c, resp)
    }
}
func WrapParamReturn[Resp any](
    svcFunc func(context.Context, string) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id := c.Params("id")
        resp, err := svcFunc(c.Context(), id)
        if err != nil {
            return ErrorResponse(c, fiber.StatusNotFound, err.Error())
        }
        return SuccessResponse(c, resp)
    }
}
func WrapNoBody[Resp any](
    svcFunc func(context.Context) (*Resp, error),
) fiber.Handler {
    return func(c *fiber.Ctx) error {
        resp, err := svcFunc(c.Context())
        if err != nil {
            return ErrorResponse(c, fiber.StatusBadRequest, err.Error())
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