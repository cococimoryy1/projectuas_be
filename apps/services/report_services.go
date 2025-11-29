package services

import (
    "context"
    "errors"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
)

type ReportService struct {
    Repo        repository.ReportRepository
    StudentRepo repository.StudentRepository
}

func NewReportService(r repository.ReportRepository, s repository.StudentRepository) *ReportService {
    return &ReportService{Repo: r, StudentRepo: s}
}

func (s *ReportService) Statistics(ctx context.Context) (*models.ReportStatisticsResponse, error) {

    claims := ctx.Value("claims").(*models.JwtCustomClaims)

    switch claims.RoleName {

    case "Admin":
        return s.Repo.StatsAll(ctx)

    case "Dosen Wali":
        return s.Repo.StatsByAdvisor(ctx, claims.LecturerID)

    case "Mahasiswa":
        return s.Repo.StatsByStudent(ctx, claims.StudentID)
    }

    return nil, errors.New("forbidden")
}

func (s *ReportService) StudentReport(ctx context.Context, id string) (*models.ReportStudentDetail, error) {

    claims := ctx.Value("claims").(*models.JwtCustomClaims)

    // mahasiswa → diri sendiri
    if claims.RoleName == "Mahasiswa" && claims.StudentID != id {
        return nil, errors.New("forbidden")
    }

    // dosen wali → advisee
    if claims.RoleName == "Dosen Wali" {
        allowed, _ := s.StudentRepo.IsAdvisorOf(ctx, claims.LecturerID, id)
        if !allowed {
            return nil, errors.New("forbidden")
        }
    }

    return s.Repo.StudentSummary(ctx, id)
}
