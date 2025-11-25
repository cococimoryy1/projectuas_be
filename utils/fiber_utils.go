package utils

import (
    "context"
    "encoding/json"
    "path/filepath"
    "strings"
    "github.com/google/uuid"

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
        studentID, ok := c.Locals("studentID").(string)
        if !ok {
            return ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated")
        }

        list, err := svcFunc(c.Context(), studentID)
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


func WrapLogicParam[Req any, Resp any](
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
func WrapCreateAchievement(
    svcFunc func(context.Context, models.CreateAchievementParsed, string) (*models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        studentID := c.Locals("studentID").(string)


        // Parse fields
        parsed := models.CreateAchievementParsed{
            Title:           c.FormValue("title"),
            Description:     c.FormValue("description"),
            AchievementType: c.FormValue("achievementType"),
        }

        // Parse dynamic fields (details json)
        detailsJson := c.FormValue("details")
        if detailsJson != "" {
            var details map[string]interface{}
            if err := json.Unmarshal([]byte(detailsJson), &details); err == nil {
                parsed.Details = details
            }
        }

        // Tags comma separated
        tags := c.FormValue("tags")
        if tags != "" {
            parsed.Tags = strings.Split(tags, ",")
        }

        // File
        file, err := c.FormFile("file")
        if err != nil {
            return ErrorResponse(c, 400, "file is required")
        }

        if file.Size > 2*1024*1024 {
            return ErrorResponse(c, 400, "file exceeds 2MB")
        }

        mime := file.Header.Get("Content-Type")
        allowed := map[string]bool{
            "application/pdf": true,
            "image/jpeg": true,
            "image/png": true,
        }
        if !allowed[mime] {
            return ErrorResponse(c, 400, "only PDF/JPG/PNG allowed")
        }

        // Save file
        ext := filepath.Ext(file.Filename)
        filename := uuid.New().String() + ext
        path := "./uploads/achievements/" + filename

        if err := c.SaveFile(file, path); err != nil {
            return ErrorResponse(c, 500, "failed saving file")
        }

        parsed.FilePath = filename
        parsed.FileType = mime

        // call service
       resp, svcErr := svcFunc(c.Context(), parsed, studentID)
        if svcErr != nil {
            return ErrorResponse(c, 400, svcErr.Error())
        }

        return SuccessResponse(c, resp)
    }
}
