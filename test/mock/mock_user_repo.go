package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockUserRepo struct {
    CreateUserFunc func(ctx context.Context, req models.CreateUserRequest) (string, error)
    GetUserByIDFunc func(ctx context.Context, id string) (*models.User, error)
}

func (m *MockUserRepo) CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
    return m.CreateUserFunc(ctx, req)
}

func (m *MockUserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    return m.GetUserByIDFunc(ctx, id)
}
