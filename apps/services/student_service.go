package services

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
)

type StudentService struct {
    Repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
    return &StudentService{Repo: repo}
}

func (s *StudentService) List(ctx context.Context) (*[]models.StudentListResponse, error) {
    result, err := s.Repo.ListStudents(ctx)
    if err != nil {
        return nil, err
    }
    return &result, nil
}
