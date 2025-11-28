package helper

// import (
// 	"BE_PROJECTUAS/utils"
// 	"BE_PROJECTUAS/apps/models"
// 	"github.com/gofiber/fiber/v2"
// 	"errors"
// 	"strings"
// )

// func AuthCheck(c *fiber.Ctx) {
//     authHeader := c.Get("Authorization")
//     if authHeader == "" {
//         return
//     }

//     parts := strings.Split(authHeader, " ")
//     if len(parts) != 2 {
//         return
//     }

//     claims, err := utils.ValidateToken(parts[1])
//     if err != nil {
//         return
//     }

//     c.Locals("claims", claims)
//     c.Locals("userID", claims.UserID)
//     c.Locals("studentID", claims.StudentID)
//     c.Locals("lecturerID", claims.LecturerID)
//     c.Locals("permissions", claims.Permissions)
// }

// func ExtractClaims(c *fiber.Ctx) (*models.JwtCustomClaims, error) {
// 	raw := c.Locals("claims")
// 	if raw == nil {
// 		return nil, errors.New("unauthorized")
// 	}
// 	return raw.(*models.JwtCustomClaims), nil
// }
