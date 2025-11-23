package repository

import (
    "BE_PROJECTUAS/apps/models"
    "context"
)

type UserRepository interface {
    FindByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error)
    GetPermissionsByRoleID(ctx context.Context, roleID string) ([]string, error)
    FindByID(ctx context.Context, id string) (*models.User, error)
}

type AchievementRepository interface {
    CreateAchievementReference(ctx context.Context, a models.Achievement) (string, error)
    UpdateStatus(ctx context.Context, id string, status string) error
    ListByStudent(ctx context.Context, studentID string) ([]models.Achievement, error)
    ListByAdvisorStudents(ctx context.Context, advisorID string) ([]models.Achievement, error) // Tambah FR-006
    // Tambah jika perlu: ListAll for admin FR-010
}