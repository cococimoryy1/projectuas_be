package tests

import (
	"context"
	"testing"

	"BE_PROJECTUAS/apps/models"
	"BE_PROJECTUAS/apps/services"
	"BE_PROJECTUAS/test/mock"
)

func TestGetStudentAchievements(t *testing.T) {

    mockStudent := &mocks.MockStudentRepo{
        GetByIDFunc: func(ctx context.Context, id string) (*models.StudentDetailResponse, error) {
            return &models.StudentDetailResponse{
                ID:        "s1",      // internal DB ID
                StudentID: "S001",    // NIM
            }, nil
        },
    }

    mockAch := &mocks.MockAchievementRepo{
        ListByStudentIDFunc: func(ctx context.Context, id string) ([]models.AchievementResponse, error) {
            return []models.AchievementResponse{
                {ID: "a1", StudentID: "s1", Status: "verified"},
            }, nil
        },
    }

    svc := services.NewStudentService(mockStudent, nil, mockAch)

    ctx := context.WithValue(context.Background(), "claims", &models.JwtCustomClaims{
        RoleName:  "Mahasiswa",
        StudentID: "s1",  // FIXED
    })

    resp, err := svc.GetStudentAchievements(ctx, "s1")
    if err != nil {
        t.Errorf("Expected success, got %v", err)
    }

    if len(resp) != 1 {
        t.Errorf("Expected 1 achievement, got %d", len(resp))
    }
}

