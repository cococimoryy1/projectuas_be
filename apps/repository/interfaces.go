package repository

import (
    "BE_PROJECTUAS/apps/models"
    "context"
)

type AuthRepository interface {
    FindByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error)
    FindByID(ctx context.Context, id string) (*models.User, error)
    GetPermissionsByRoleID(ctx context.Context, roleID string) ([]string, error)
    GetStudentByUserID(ctx context.Context, userID string) (string, error)
}

type UserRepository interface {
    ListUsers(ctx context.Context) ([]models.User, error)
    GetUserByID(ctx context.Context, id string) (*models.User, error)
    CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error)
    UpdateUser(ctx context.Context, id string, req models.UpdateUserRequest) error
    DeleteUser(ctx context.Context, id string) error
    UpdateUserRole(ctx context.Context, id string, roleID string) error
}


type AchievementRepository interface {
    CreateAchievementReference(ctx context.Context, a models.Achievement) (string, error)
    UpdateStatus(ctx context.Context, id string, status string) error
    ListByStudent(ctx context.Context, studentID string) ([]models.Achievement, error)
    ListByAdvisorStudents(ctx context.Context, advisorID string) ([]models.Achievement, error) // Tambah FR-006
    // Tambah jika perlu: ListAll for admin FR-010
    InsertMongoAchievement(ctx context.Context, doc models.AchievementMongo) (string, error)
}