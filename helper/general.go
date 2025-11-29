package helper

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "github.com/gofiber/fiber/v2"
)

func WrapLogic[Req any, Resp any](
    svcFunc func(context.Context, Req) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        body := c.Locals("parsed_body")
        req, ok := body.(Req)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        resp, err := svcFunc(c.Context(), req)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


func WrapParam(
    svcFunc func(context.Context, string) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        if err := svcFunc(c.Context(), id); err != nil {
            return c.Status(409).JSON(Error(409, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "status": "updated",
        }))
    }
}

func WrapReject(
    svcFunc func(context.Context, string, string) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        body := c.Locals("parsed_body")
        req, ok := body.(models.RejectRequest)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        if err := svcFunc(c.Context(), id, req.Note); err != nil {
            return c.Status(422).JSON(Error(422, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "status": "rejected",
        }))
    }
}

// WrapRefresh: For refresh (body with refreshToken)
func WrapListAll(
    svcFunc func(context.Context) ([]models.Achievement, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        list, err := svcFunc(c.Context())
        if err != nil {
            return c.Status(500).JSON(Error(500, err.Error()))
        }

        return c.JSON(Success(list))
    }
}

func WrapUpdate[Req any](
    svc func(context.Context, string, Req) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        body := c.Locals("parsed_body")
        req, ok := body.(Req)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        if err := svc(c.Context(), id, req); err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "status": "updated",
        }))
    }
}

func WrapLogicParam[Req any, Resp any](
    svcFunc func(context.Context, string, Req) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        body := c.Locals("parsed_body")
        req, ok := body.(Req)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        resp, err := svcFunc(c.Context(), id, req)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}

func WrapParamResp[Resp any](
    svcFunc func(context.Context, string) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        resp, err := svcFunc(c.Context(), id)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


func WrapUpdateResp[Req any, Resp any](
    svcFunc func(context.Context, string, Req) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        body := c.Locals("parsed_body")
        req, ok := body.(Req)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        resp, err := svcFunc(c.Context(), id, req)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


func WrapParamReturn[Resp any](
    svcFunc func(context.Context, string) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        // inject fiber.Ctx ke context
        ctx := context.WithValue(c.Context(), "fiberCtx", c)

        resp, err := svcFunc(ctx, id)
        if err != nil {
            return c.Status(404).JSON(Error(404, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


func WrapNoBody[Resp any](
    svcFunc func(context.Context) (*Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        resp, err := svcFunc(c.Context())
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}

func WrapProfile(
    svcFunc func(context.Context, string) (*models.UserResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        userID, ok := c.Locals("userID").(string)
        if !ok {
            return c.Status(401).JSON(Error(401, "user not authenticated"))
        }

        resp, err := svcFunc(c.Context(), userID)
        if err != nil {
            return c.Status(404).JSON(Error(404, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


func WrapLogout(
    svcFunc func(context.Context) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        if err := svcFunc(c.Context()); err != nil {
            return c.Status(500).JSON(Error(500, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "message": "Logged out successfully",
        }))
    }
}

func WrapRefresh(
    svcFunc func(context.Context, string) (*models.LoginResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        body := c.Locals("parsed_body")

        req, ok := body.(models.RefreshRequest)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        if req.RefreshToken == "" {
            return c.Status(400).JSON(Error(400, "refreshToken missing"))
        }

        resp, err := svcFunc(c.Context(), req.RefreshToken)
        if err != nil {
            return c.Status(401).JSON(Error(401, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}


