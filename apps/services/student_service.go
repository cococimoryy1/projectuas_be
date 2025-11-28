package services

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    "errors"
    // "github.com/gofiber/fiber/v2"
)

type StudentService struct {
    Repo     repository.StudentRepository
    AuthRepo repository.AuthRepository
    AchRepo  repository.AchievementRepository
}

func NewStudentService( repo repository.StudentRepository, auth repository.AuthRepository, ach repository.AchievementRepository,) *StudentService {
    return &StudentService{
        Repo:     repo,
        AuthRepo: auth,
        AchRepo:  ach,
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

func (s *StudentService) GetStudentAchievements(ctx context.Context, id string) ([]models.AchievementResponse, error) {

    claims := ctx.Value("claims").(*models.JwtCustomClaims)

    // Ambil student detail
    student, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // Mahasiswa hanya boleh lihat miliknya
    if claims.RoleName == "Mahasiswa" && claims.StudentID != student.ID {
        return nil, errors.New("forbidden")
    }

    // Dosen wali hanya boleh lihat advisee-nya
    if claims.RoleName == "Dosen Wali" {
        allowed, _ := s.Repo.IsAdvisorOf(ctx, claims.LecturerID, id)
        if !allowed {
            return nil, errors.New("forbidden")
        }
    }

    // Admin boleh semua
    return s.AchRepo.ListByStudentID(ctx, student.ID)
}

