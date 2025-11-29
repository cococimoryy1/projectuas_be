package helper
import (
    "context"
    "github.com/gofiber/fiber/v2"
)


func WrapParamReturnList[Resp any](
    svcFunc func(context.Context, string) ([]Resp, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")
        ctx := context.WithValue(c.Context(), "claims", c.Locals("claims"))

        resp, err := svcFunc(ctx, id)
        if err != nil {
            return c.Status(404).JSON(Error(404, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}
func WrapParamBody[Req any](
    svcFunc func(context.Context, string, Req) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        // Parse body
        bodyRaw := c.Locals("parsed_body")
        req, ok := bodyRaw.(Req)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        ctx := context.WithValue(c.Context(), "fiberCtx", c)

        err := svcFunc(ctx, id, req)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success("updated"))
    }
}
