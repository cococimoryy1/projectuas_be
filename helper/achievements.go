package wrappers

import (
    "context"
    "encoding/json"
    "path/filepath"
    "strings"
    "time"
    "github.com/google/uuid"
    "github.com/gofiber/fiber/v2"

    "BE_PROJECTUAS/apps/models"
)


// --- Create Achievement (form-data + file upload) ---
func WrapCreateAchievement(
    svc func(context.Context, models.CreateAchievementParsed, string) (*models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        studentID := c.Locals("studentID").(string)

        parsed := models.CreateAchievementParsed{
            Title: c.FormValue("title"),
            Description: c.FormValue("description"),
            AchievementType: c.FormValue("achievementType"),
        }

        if details := c.FormValue("details"); details != "" {
            json.Unmarshal([]byte(details), &parsed.Details)
        }

        if tags := c.FormValue("tags"); tags != "" {
            parsed.Tags = strings.Split(tags, ",")
        }

        file, err := c.FormFile("file")
        if err != nil {
            return ErrorResponse(c, 400, "file required")
        }

        filename := uuid.New().String() + filepath.Ext(file.Filename)
        path := "./uploads/achievements/" + filename

        if err := c.SaveFile(file, path); err != nil {
            return ErrorResponse(c, 500, "failed saving file")
        }

        parsed.FilePath = filename
        parsed.FileType = file.Header.Get("Content-Type")

        resp, svcErr := svc(c.Context(), parsed, studentID)
        if svcErr != nil {
            return ErrorResponse(c, 400, svcErr.Error())
        }

        return SuccessResponse(c, resp)
    }
}


// --- List Student's Own Achievement ---
func WrapListStudent(
    svc func(context.Context, string) ([]models.Achievement, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {
        studentID := c.Locals("studentID").(string)

        list, err := svc(c.Context(), studentID)
        if err != nil {
            return ErrorResponse(c, 500, err.Error())
        }
        return SuccessResponse(c, list)
    }
}


// --- List Advisor's Advisees' Achievements ---
func WrapListAdvisor(
    svc func(context.Context, string) ([]models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {
        advisorID := c.Locals("userID").(string)

        list, err := svc(c.Context(), advisorID)
        if err != nil {
            return ErrorResponse(c, 500, err.Error())
        }

        return SuccessResponse(c, list)
    }
}

func WrapUpdateDraft(
    svc func(context.Context, string, models.UpdateAchievementRequest) (*models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {
        id := c.Params("id")

        req, err := ParseBody[models.UpdateAchievementRequest](c)
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
func WrapDeleteDraft(
    svc func(context.Context, string, string) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        userID := c.Locals("userID").(string)

        if err := svc(c.Context(), id, userID); err != nil {
            return ErrorResponse(c, 400, err.Error())
        }

        return SuccessResponse(c, fiber.Map{
            "message": "draft deleted successfully",
        })
    }
}
func WrapUploadAttachment(
    svc func(context.Context, string, models.AttachmentMongo) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        file, err := c.FormFile("file")
        if err != nil {
            return ErrorResponse(c, 400, "file required")
        }

        filename := uuid.New().String() + filepath.Ext(file.Filename)
        path := "./uploads/attachments/" + filename
        c.SaveFile(file, path)

        att := models.AttachmentMongo{
            FileName:   filename,
            FileUrl:    "/uploads/attachments/" + filename,
            FileType:   file.Header.Get("Content-Type"),
            UploadedAt: time.Now(),
        }

        if err := svc(c.Context(), id, att); err != nil {
            return ErrorResponse(c, 400, err.Error())
        }

        return SuccessResponse(c, fiber.Map{"message": "attachment uploaded"})
    }
}
