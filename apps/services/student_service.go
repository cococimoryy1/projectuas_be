package services

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    // "github.com/gofiber/fiber/v2"
)

type StudentService struct {
    Repo     repository.StudentRepository
    AuthRepo repository.AuthRepository
}

func NewStudentService(repo repository.StudentRepository, auth repository.AuthRepository) *StudentService {
    return &StudentService{
        Repo:     repo,
        AuthRepo: auth,
    }
}

func (s *StudentService) List(ctx context.Context) (*[]models.StudentListResponse, error) {
    list, err := s.Repo.ListStudents(ctx)
    if err != nil {
        return nil, err
    }
    return &list, nil
}

func (s *StudentService) GetByID(ctx context.Context, id string) (*models.StudentDetailResponse, error) {
    return s.Repo.GetByID(ctx, id)
}
