package services

import (
    "context"
	"errors"

    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
)

type LecturerService struct {
    Repo      repository.LecturerRepository
    StudentRepo repository.StudentRepository
}

func NewLecturerService(repo repository.LecturerRepository, srepo repository.StudentRepository) *LecturerService {
    return &LecturerService{
        Repo: repo,
        StudentRepo: srepo,
    }
}


func (s *LecturerService) List(ctx context.Context) (*[]models.LecturerListResponse, error) {

    claims := ctx.Value("claims").(*models.JwtCustomClaims)

    // ADMIN → semua
    if claims.RoleName == "Admin" {
        list, err := s.Repo.ListLecturers(ctx)
        return &list, err
    }

    // DOSEN WALI → hanya dirinya sendiri
    if claims.RoleName == "Dosen Wali" {
        data, err := s.Repo.GetByID(ctx, claims.LecturerID)
        if err != nil {
            return nil, err
        }
        list := []models.LecturerListResponse{
            {
                ID:         data.ID,
                LecturerID: data.LecturerID,
                FullName:   data.FullName,
                Email:      data.Email,
                Department: data.Department,
                AdviseeCount: 0,
            },
        }
        return &list, nil
    }

    // MAHASISWA → hanya dosen walinya
    if claims.RoleName == "Mahasiswa" {

        advisorID, err := s.Repo.GetAdvisorByStudentID(ctx, claims.StudentID)
        if err != nil {
            return nil, err
        }

        data, err := s.Repo.GetByID(ctx, advisorID)
        if err != nil {
            return nil, err
        }

        list := []models.LecturerListResponse{
            {
                ID:         data.ID,
                LecturerID: data.LecturerID,
                FullName:   data.FullName,
                Email:      data.Email,
                Department: data.Department,
            },
        }

        return &list, nil
    }

    return nil, errors.New("forbidden")
}
func (s *LecturerService) GetAdvisees(ctx context.Context, lecturerID string) (*[]models.StudentListResponse, error) {

    claims := ctx.Value("claims").(*models.JwtCustomClaims)

    // ADMIN → semua lecturer bebas
    if claims.RoleName == "Admin" {
        list, err := s.Repo.ListAdvisees(ctx, lecturerID)
        return &list, err
    }

    // DOSEN WALI → hanya datanya sendiri
    if claims.RoleName == "Dosen Wali" && claims.LecturerID != lecturerID {
        return nil, errors.New("forbidden")
    }

    // MAHASISWA → hanya boleh lihat advises dari dosen walinya sendiri
    if claims.RoleName == "Mahasiswa" {

        advisorID, err := s.Repo.GetAdvisorByStudentID(ctx, claims.StudentID)
        if err != nil {
            return nil, err
        }

        if advisorID != lecturerID {
            return nil, errors.New("forbidden")
        }
    }

    // EXECUTE QUERY
    list, err := s.Repo.ListAdvisees(ctx, lecturerID)
    if err != nil {
        return nil, err
    }

    return &list, nil
}


