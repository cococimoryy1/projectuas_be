package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockLecturerRepo struct {
    ListFunc       func(ctx context.Context) ([]models.LecturerListResponse, error)
    GetByIDFunc    func(ctx context.Context, id string) (*models.LecturerDetailResponse, error)
    ListAdviseesFunc func(ctx context.Context, id string) ([]models.StudentListResponse, error)
}

func (m *MockLecturerRepo) ListLecturers(ctx context.Context) ([]models.LecturerListResponse, error) {
    return m.ListFunc(ctx)
}

func (m *MockLecturerRepo) GetByID(ctx context.Context, id string) (*models.LecturerDetailResponse, error) {
    return m.GetByIDFunc(ctx, id)
}

func (m *MockLecturerRepo) ListAdvisees(ctx context.Context, id string) ([]models.StudentListResponse, error) {
    return m.ListAdviseesFunc(ctx, id)
}
