package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockAchievementRepo struct {
    CreateAchievementReferenceFunc func(context.Context, models.Achievement) (string, error)
    UpdateStatusFunc               func(context.Context, string, string) error
    ListByStudentFunc              func(context.Context, string) ([]models.Achievement, error)
    ListByAdvisorStudentsFunc      func(context.Context, string) ([]models.Achievement, error)
    InsertMongoAchievementFunc     func(context.Context, models.AchievementMongo) (string, error)
    GetByIDFunc                    func(context.Context, string) (*models.Achievement, error)
    VerifyAchievementFunc          func(context.Context, string, string) error
    GetMongoByIDFunc               func(context.Context, string) (*models.AchievementMongo, error)
    RejectAchievementFunc          func(context.Context, string, string, string) error
    IsAdvisorOfFunc                func(context.Context, string, string) (bool, error)
    SubmitAchievementFunc          func(context.Context, string) error
    UpdateMongoAchievementFunc     func(context.Context, string, models.UpdateAchievementRequest) error
    TouchUpdatedAtFunc             func(context.Context, string) error
    SoftDeleteFunc                 func(context.Context, string, string) error
    SoftDeleteMongoFunc            func(context.Context, string) error
    AddAttachmentFunc              func(context.Context, string, models.AttachmentMongo) error
    ListByStudentIDFunc            func(context.Context, string) ([]models.AchievementResponse, error)
}

// ==== Implementasi Interface ====

func (m *MockAchievementRepo) CreateAchievementReference(ctx context.Context, a models.Achievement) (string, error) {
    return m.CreateAchievementReferenceFunc(ctx, a)
}

func (m *MockAchievementRepo) UpdateStatus(ctx context.Context, id string, status string) error {
    return m.UpdateStatusFunc(ctx, id, status)
}

func (m *MockAchievementRepo) ListByStudent(ctx context.Context, id string) ([]models.Achievement, error) {
    return m.ListByStudentFunc(ctx, id)
}

func (m *MockAchievementRepo) ListByAdvisorStudents(ctx context.Context, id string) ([]models.Achievement, error) {
    return m.ListByAdvisorStudentsFunc(ctx, id)
}

func (m *MockAchievementRepo) InsertMongoAchievement(ctx context.Context, doc models.AchievementMongo) (string, error) {
    return m.InsertMongoAchievementFunc(ctx, doc)
}

func (m *MockAchievementRepo) GetByID(ctx context.Context, id string) (*models.Achievement, error) {
    return m.GetByIDFunc(ctx, id)
}

func (m *MockAchievementRepo) VerifyAchievement(ctx context.Context, id, lecturerID string) error {
    return m.VerifyAchievementFunc(ctx, id, lecturerID)
}

func (m *MockAchievementRepo) GetMongoByID(ctx context.Context, id string) (*models.AchievementMongo, error) {
    return m.GetMongoByIDFunc(ctx, id)
}

func (m *MockAchievementRepo) RejectAchievement(ctx context.Context, id, lecturerID, note string) error {
    return m.RejectAchievementFunc(ctx, id, lecturerID, note)
}

func (m *MockAchievementRepo) IsAdvisorOf(ctx context.Context, lecturerID, studentID string) (bool, error) {
    return m.IsAdvisorOfFunc(ctx, lecturerID, studentID)
}

func (m *MockAchievementRepo) SubmitAchievement(ctx context.Context, id string) error {
    return m.SubmitAchievementFunc(ctx, id)
}

func (m *MockAchievementRepo) UpdateMongoAchievement(ctx context.Context, id string, req models.UpdateAchievementRequest) error {
    return m.UpdateMongoAchievementFunc(ctx, id, req)
}

func (m *MockAchievementRepo) TouchUpdatedAt(ctx context.Context, id string) error {
    return m.TouchUpdatedAtFunc(ctx, id)
}

func (m *MockAchievementRepo) SoftDelete(ctx context.Context, id string, userID string) error {
    return m.SoftDeleteFunc(ctx, id, userID)
}

func (m *MockAchievementRepo) SoftDeleteMongo(ctx context.Context, id string) error {
    return m.SoftDeleteMongoFunc(ctx, id)
}

func (m *MockAchievementRepo) AddAttachment(ctx context.Context, id string, att models.AttachmentMongo) error {
    return m.AddAttachmentFunc(ctx, id, att)
}

func (m *MockAchievementRepo) ListByStudentID(ctx context.Context, studentID string) ([]models.AchievementResponse, error) {
    return m.ListByStudentIDFunc(ctx, studentID)
}
