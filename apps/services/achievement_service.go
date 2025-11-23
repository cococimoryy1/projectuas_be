package services

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    "context"
    "errors"
    // "time"

    "github.com/google/uuid"
)

type AchievementService struct {
    Repo repository.AchievementRepository
}

func NewAchievementService(repo repository.AchievementRepository) *AchievementService {
    return &AchievementService{Repo: repo}
}

func (s *AchievementService) Create(ctx context.Context, req models.CreateAchievementRequest, studentID string) (*models.AchievementResponse, error) {
    if req.Title == "" {
        return nil, errors.New("title required")
    }

    // Placeholder Mongo insert (SRS hal.6-7)
    // mongoID, err := mongoRepo.Insert(ctx, req) // Implement: {studentId: studentID, title: req.Title, details: req.Details}
    mongoID := uuid.New().String() // Placeholder

    ach := models.Achievement{
        ID:        uuid.New().String(),
        StudentID: studentID,
        MongoID:   mongoID,
        Status:    "draft",
    }

    refID, err := s.Repo.CreateAchievementReference(ctx, ach)
    if err != nil {
        return nil, errors.New("failed to create achievement")
    }

    return &models.AchievementResponse{ID: refID, Status: "draft"}, nil
}

func (s *AchievementService) Submit(ctx context.Context, id string) error {
    // Check status draft if needed
    return s.Repo.UpdateStatus(ctx, id, "submitted")
}

func (s *AchievementService) Verify(ctx context.Context, id string) error {
    return s.Repo.UpdateStatus(ctx, id, "verified")
}

func (s *AchievementService) Reject(ctx context.Context, id string, note string) error {
    // Extend repo for note if needed (SRS hal.5)
    return s.Repo.UpdateStatus(ctx, id, "rejected")
}

func (s *AchievementService) ListForStudent(ctx context.Context, studentID string) ([]models.Achievement, error) {
    return s.Repo.ListByStudent(ctx, studentID)
}

func (s *AchievementService) ListForAdvisor(ctx context.Context, advisorID string) ([]models.Achievement, error) {
    return s.Repo.ListByAdvisorStudents(ctx, advisorID)
}

// Update, Delete, History, Upload (placeholder)
func (s *AchievementService) Update(ctx context.Context, id string, req models.CreateAchievementRequest) error {
    // Placeholder FR-003 update draft (check status, update Mongo + Postgres)
    return errors.New("not implemented") // Expand: Update Mongo by mongoID, then Postgres if needed
}

func (s *AchievementService) Delete(ctx context.Context, id string) error {
    return errors.New("not implemented") // FR-005 soft delete
}

func (s *AchievementService) GetHistory(ctx context.Context, id string) error {
    return errors.New("not implemented") // SRS hal.11
}

func (s *AchievementService) UploadAttachment(ctx context.Context, id string) error {
    return errors.New("not implemented") // FR-003 attachments
}