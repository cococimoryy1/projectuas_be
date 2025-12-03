package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockAuthRepo struct {
    FindByUsernameOrEmailFunc func(ctx context.Context, identifier string) (*models.User, error)
    FindByIDFunc              func(ctx context.Context, id string) (*models.User, error)
    GetPermissionsByRoleIDFunc func(ctx context.Context, roleID string) ([]string, error)
    GetStudentByUserIDFunc     func(ctx context.Context, userID string) (string, error)
    GetLecturerByUserIDFunc    func(ctx context.Context, userID string) (string, error)
}

func (m *MockAuthRepo) FindByUsernameOrEmail(ctx context.Context, id string) (*models.User, error) {
    return m.FindByUsernameOrEmailFunc(ctx, id)
}
func (m *MockAuthRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
    return m.FindByIDFunc(ctx, id)
}
func (m *MockAuthRepo) GetPermissionsByRoleID(ctx context.Context, roleID string) ([]string, error) {
    return m.GetPermissionsByRoleIDFunc(ctx, roleID)
}
func (m *MockAuthRepo) GetStudentByUserID(ctx context.Context, userID string) (string, error) {
    return m.GetStudentByUserIDFunc(ctx, userID)
}
func (m *MockAuthRepo) GetLecturerByUserID(ctx context.Context, userID string) (string, error) {
    return m.GetLecturerByUserIDFunc(ctx, userID)
}
