package tests

import (
    "context"
    "errors"
    "testing"
	"os"

    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/services"
    mocks "BE_PROJECTUAS/test/mock"
)

func TestAuthLoginSuccess(t *testing.T) {

    // WAJIB: agar Login tidak gagal karena missing secret
    os.Setenv("JWT_SECRET", "testsecret")

    mockRepo := &mocks.MockAuthRepo{
        FindByUsernameOrEmailFunc: func(ctx context.Context, u string) (*models.User, error) {
            return &models.User{
                ID:           "123",
                Username:     "admin",
                Email:        "admin@test.com",
                PasswordHash: "$2a$12$Xo6Jdr4eyIWfW1StlZ0tlexncyn6WlziH5qvUtkOOnXW9m5dqVWYa", // hash 12345678
                RoleID:       "admin-role",
                RoleName:     "Admin",
                IsActive:     true,
            }, nil
        },

        GetPermissionsByRoleIDFunc: func(ctx context.Context, roleID string) ([]string, error) {
            return []string{"user:manage"}, nil
        },

        GetStudentByUserIDFunc: func(ctx context.Context, userID string) (string, error) {
            return "", nil
        },

        GetLecturerByUserIDFunc: func(ctx context.Context, userID string) (string, error) {
            return "", nil
        },
    }

    svc := services.NewAuthService(mockRepo)

    req := models.LoginRequest{
        Username: "admin",
        Password: "12345678",
    }

    _, err := svc.Login(context.Background(), req)
    if err != nil {
        t.Errorf("Expected success, got: %v", err)
    }
}


func TestAuthLoginInvalidUser(t *testing.T) {
    mockRepo := &mocks.MockAuthRepo{
        FindByUsernameOrEmailFunc: func(ctx context.Context, u string) (*models.User, error) {
            return nil, errors.New("not found")
        },
    }

    svc := services.NewAuthService(mockRepo)

    req := models.LoginRequest{
        Username: "unknown",
        Password: "123456",
    }

    _, err := svc.Login(context.Background(), req)
    if err == nil {
        t.Error("Expected error, got nil")
    }
}
