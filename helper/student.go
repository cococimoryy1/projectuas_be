package helper

// import "github.com/gofiber/fiber/v2"

// func WrapParamReturn(
//     svc func(ctx context.Context, id string) (interface{}, error),
// ) fiber.Handler {

//     return func(c *fiber.Ctx) error {

//         id := c.Params("id")

//         result, err := svc(c.Context(), id)
//         if err != nil {
//             return c.Status(500).JSON(Error(500, err.Error()))
//         }

//         return c.JSON(Success(result))
//     }
// }
