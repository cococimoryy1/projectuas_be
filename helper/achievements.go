package helper

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


// ==========================================================
// CREATE ACHIEVEMENT (multipart/form-data)
// ==========================================================
func WrapCreateAchievement(
    svc func(context.Context, models.CreateAchievementParsed, string) (*models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        studentID := c.Locals("studentID").(string)

        parsed := models.CreateAchievementParsed{
            Title:           c.FormValue("title"),
            Description:     c.FormValue("description"),
            AchievementType: c.FormValue("achievementType"),
        }

        // Parse JSON string jika ada
        if details := c.FormValue("details"); details != "" {
            json.Unmarshal([]byte(details), &parsed.Details)
        }

        if tags := c.FormValue("tags"); tags != "" {
            parsed.Tags = strings.Split(tags, ",")
        }

        // File upload
        file, err := c.FormFile("file")
        if err != nil {
            return c.Status(400).JSON(Error(400, "file required"))
        }

        filename := uuid.New().String() + filepath.Ext(file.Filename)
        path := "./uploads/achievements/" + filename

        if err := c.SaveFile(file, path); err != nil {
            return c.Status(500).JSON(Error(500, "failed saving file"))
        }

        parsed.FilePath = filename
        parsed.FileType = file.Header.Get("Content-Type")

        // Call service
        resp, svcErr := svc(c.Context(), parsed, studentID)
        if svcErr != nil {
            return c.Status(400).JSON(Error(400, svcErr.Error()))
        }

        return c.JSON(Success(resp))
    }
}



// ==========================================================
// LIST STUDENT'S OWN ACHIEVEMENTS
// ==========================================================
func WrapListStudent(
    svc func(context.Context, string) ([]models.Achievement, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        studentID := c.Locals("studentID").(string)

        list, err := svc(c.Context(), studentID)
        if err != nil {
            return c.Status(500).JSON(Error(500, err.Error()))
        }

        return c.JSON(Success(list))
    }
}



// ==========================================================
// LIST ADVISOR'S ADVISEE ACHIEVEMENTS
// ==========================================================
func WrapListAdvisor(
    svc func(context.Context, string) ([]models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        advisorID := c.Locals("userID").(string)

        list, err := svc(c.Context(), advisorID)
        if err != nil {
            return c.Status(500).JSON(Error(500, err.Error()))
        }

        return c.JSON(Success(list))
    }
}



// ==========================================================
// UPDATE DRAFT ACHIEVEMENT (JSON)
// ==========================================================
func WrapUpdateDraft(
    svc func(context.Context, string, models.UpdateAchievementRequest) (*models.AchievementResponse, error),
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        body := c.Locals("parsed_body")
        req, ok := body.(models.UpdateAchievementRequest)
        if !ok {
            return c.Status(400).JSON(Error(400, "invalid parsed body"))
        }

        resp, err := svc(c.Context(), id, req)
        if err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(resp))
    }
}



// ==========================================================
// DELETE DRAFT ACHIEVEMENT
// ==========================================================
func WrapDeleteDraft(
    svc func(context.Context, string, string) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")
        userID := c.Locals("userID").(string)

        if err := svc(c.Context(), id, userID); err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "message": "draft deleted successfully",
        }))
    }
}



// ==========================================================
// UPLOAD ATTACHMENT (multipart/form-data)
// ==========================================================
func WrapUploadAttachment(
    svc func(context.Context, string, models.AttachmentMongo) error,
) fiber.Handler {

    return func(c *fiber.Ctx) error {

        id := c.Params("id")

        file, err := c.FormFile("file")
        if err != nil {
            return c.Status(400).JSON(Error(400, "file required"))
        }

        filename := uuid.New().String() + filepath.Ext(file.Filename)
        path := "./uploads/attachments/" + filename

        if err := c.SaveFile(file, path); err != nil {
            return c.Status(500).JSON(Error(500, "failed saving attachment"))
        }

        att := models.AttachmentMongo{
            FileName:   filename,
            FileUrl:    "/uploads/attachments/" + filename,
            FileType:   file.Header.Get("Content-Type"),
            UploadedAt: time.Now(),
        }

        if err := svc(c.Context(), id, att); err != nil {
            return c.Status(400).JSON(Error(400, err.Error()))
        }

        return c.JSON(Success(map[string]any{
            "message": "attachment uploaded",
        }))
    }
}
