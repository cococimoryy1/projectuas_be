package tests

import (
	"context"
	"testing"

	"BE_PROJECTUAS/apps/models"
	"BE_PROJECTUAS/apps/services"
	"BE_PROJECTUAS/test/mock"
)

func TestReportStatisticsAdmin(t *testing.T) {
	mockRepo := &mocks.MockReportRepo{
		StatsAllFunc: func(ctx context.Context) (*models.ReportStatisticsResponse, error) {
			return &models.ReportStatisticsResponse{
				TypeStats: []models.AchievementTypeStat{
					{Type: "competition", Total: 5},
				},
			}, nil
		},
	}

	svc := services.NewReportService(mockRepo, nil)

	ctx := context.WithValue(context.Background(), "claims", &models.JwtCustomClaims{
		RoleName: "Admin",
	})

	resp, err := svc.Statistics(ctx)
	if err != nil {
		t.Errorf("Expected success, got %v", err)
	}
	if len(resp.TypeStats) != 1 {
		t.Errorf("Expected 1 stat result")
	}
}

func TestStudentReport(t *testing.T) {
	mockRepo := &mocks.MockReportRepo{
		StudentSummaryFunc: func(ctx context.Context, id string) (*models.ReportStudentDetail, error) {
			return &models.ReportStudentDetail{
				StudentID:        "S01",
				FullName:         "Budi",
				TotalAchievements: 3,
			}, nil
		},
	}

	svc := services.NewReportService(mockRepo, nil)

	ctx := context.WithValue(context.Background(), "claims", &models.JwtCustomClaims{
		RoleName:  "Mahasiswa",
		StudentID: "S01",
	})

	_, err := svc.StudentReport(ctx, "S01")
	if err != nil {
		t.Errorf("Expected success but got %v", err)
	}
}
