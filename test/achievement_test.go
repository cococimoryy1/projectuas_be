package tests

import (
	"context"
	"testing"

	"BE_PROJECTUAS/apps/models"
	"BE_PROJECTUAS/apps/services"
	"BE_PROJECTUAS/test/mock"
)

func TestCreateAchievement(t *testing.T) {

    mockRepo := &mocks.MockAchievementRepo{
        CreateAchievementReferenceFunc: func(ctx context.Context, a models.Achievement) (string, error) {
            return "new-ach-id", nil
        },
        InsertMongoAchievementFunc: func(ctx context.Context, doc models.AchievementMongo) (string, error) {
            return "mongo-123", nil
        },

        // stub required but unused methods
        UpdateMongoAchievementFunc: func(ctx context.Context, id string, r models.UpdateAchievementRequest) error { return nil },
        TouchUpdatedAtFunc: func(ctx context.Context, id string) error { return nil },
        GetByIDFunc: func(ctx context.Context, id string) (*models.Achievement, error) { return nil, nil },
    }

    svc := services.NewAchievementService(mockRepo)

    req := models.CreateAchievementParsed{
        Title: "Prestasi UI/UX",
        Description: "Testing",
        AchievementType: "competition",
        FilePath: "proof.png",
        FileType: "image/png",
    }

    _, err := svc.Create(context.Background(), req, "student-123")
    if err != nil {
        t.Errorf("Expected success, got: %v", err)
    }
}


func TestSubmitAchievement(t *testing.T) {

    mockRepo := &mocks.MockAchievementRepo{

        // WAJIB! Karena service memanggil GetByID
        GetByIDFunc: func(ctx context.Context, id string) (*models.Achievement, error) {
            return &models.Achievement{
                ID:        id,
                StudentID: "student-123",
                Status:    "draft",
            }, nil
        },

        SubmitAchievementFunc: func(ctx context.Context, id string) error {
            return nil
        },

        TouchUpdatedAtFunc: func(ctx context.Context, id string) error {
            return nil
        },
    }

    svc := services.NewAchievementService(mockRepo)

    err := svc.Submit(context.Background(), "ach-001")
    if err != nil {
        t.Errorf("Expected submit OK, got: %v", err)
    }
}
