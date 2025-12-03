package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockStudentRepo struct {
    ListStudentsFunc func(ctx context.Context) ([]models.StudentListResponse, error)
    GetByIDFunc      func(ctx context.Context, id string) (*models.StudentDetailResponse, error)
    IsAdvisorOfFunc  func(ctx context.Context, advisorID string, studentID string) (bool, error)
    UpdateAdvisorFunc func(ctx context.Context, studentID string, advisorID string) error
}

func (m *MockStudentRepo) ListStudents(ctx context.Context) ([]models.StudentListResponse, error) {
    if m.ListStudentsFunc == nil {
        return []models.StudentListResponse{}, nil
    }
    return m.ListStudentsFunc(ctx)
}

func (m *MockStudentRepo) GetByID(ctx context.Context, id string) (*models.StudentDetailResponse, error) {
    return m.GetByIDFunc(ctx, id)
}

func (m *MockStudentRepo) IsAdvisorOf(ctx context.Context, advisorID string, studentID string) (bool, error) {
    if m.IsAdvisorOfFunc == nil {
        return true, nil
    }
    return m.IsAdvisorOfFunc(ctx, advisorID, studentID)
}

func (m *MockStudentRepo) UpdateAdvisor(ctx context.Context, studentID string, advisorID string) error {
    if m.UpdateAdvisorFunc == nil {
        return nil
    }
    return m.UpdateAdvisorFunc(ctx, studentID, advisorID)
}
