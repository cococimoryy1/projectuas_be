package services

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    "context"
    "errors"
    "time"

    "github.com/google/uuid"
)

type AchievementService struct {
    Repo repository.AchievementRepository
}

func NewAchievementService(repo repository.AchievementRepository) *AchievementService {
    return &AchievementService{Repo: repo}
}

func (s *AchievementService) Create(ctx context.Context, req models.CreateAchievementParsed, studentID string) (*models.AchievementResponse, error) {

    // Build mongo document
    achMongo := models.AchievementMongo{
        StudentID:       studentID,
        Title:           req.Title,
        Description:     req.Description,
        AchievementType: req.AchievementType,
        Details:         req.Details,
        Tags:            req.Tags,
        Attachments: []models.AttachmentMongo{
            {
                FileName:   req.FilePath,
                FileUrl:    "/uploads/achievements/" + req.FilePath,
                FileType:   req.FileType,
                UploadedAt: time.Now(),
            },
        },
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    mongoID, err := s.Repo.InsertMongoAchievement(ctx, achMongo)
    if err != nil {
        return nil, err
    }

    // Create reference SQL
    ref := models.Achievement{
        ID:                 uuid.New().String(),
        StudentID:          studentID,
        MongoAchievementID: mongoID,
        Status:             "draft",
    }

    refID, err := s.Repo.CreateAchievementReference(ctx, ref)
    if err != nil {
        return nil, err
    }

    return &models.AchievementResponse{
        ID:        refID,
        MongoID:   mongoID,
        Status:    "draft",
    }, nil
}

func (s *AchievementService) Update(ctx context.Context, id string, req models.UpdateAchievementRequest) (*models.AchievementResponse, error) {

    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    if ref.Status != "draft" {
        return nil, errors.New("achievement can only be updated while in draft status")
    }

    // Update Mongo
    err = s.Repo.UpdateMongoAchievement(ctx, ref.MongoAchievementID, req)
    if err != nil {
        return nil, err
    }

    // Update SQL
    err = s.Repo.TouchUpdatedAt(ctx, id)
    if err != nil {
        return nil, err
    }

    // Build response from SQL + Mongo
    return &models.AchievementResponse{
        ID:          ref.ID,
        MongoID:     ref.MongoAchievementID,
        StudentID:   ref.StudentID,
        Title:       req.Title,
        Description: req.Description,
        Category:    ref.Status,
        Status:      ref.Status,
        CreatedAt:   ref.CreatedAt.Format(time.RFC3339),
    }, nil
}


func (s *AchievementService) Submit(ctx context.Context, id string) error {

    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    if ref.Status != "draft" {
        return errors.New("only draft achievement can be submitted")
    }

    return s.Repo.SubmitAchievement(ctx, id)
}


func (s *AchievementService) Verify(ctx context.Context, id string) error {
    // Ambil data
    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    // Validasi status
    if ref.Status != "submitted" {
        return errors.New("only submitted achievements can be verified")
    }

    // VALIDASI DOSEN WALI
    lecturerID := ctx.Value("lecturerID").(string)

    allowed, err := s.Repo.IsAdvisorOf(ctx, lecturerID, ref.StudentID)
    if err != nil {
        return err
    }
    if !allowed {
        return errors.New("forbidden: not advisor of this student")
    }

    // SIMPAN USERID (FK KE USERS.ID)
    userID := ctx.Value("userID").(string)

    // UPDATE
    return s.Repo.VerifyAchievement(ctx, id, userID)
}


func (s *AchievementService) Delete(ctx context.Context, id string, userID string) error {

    // 1. Ambil prestasi
    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    // 2. Pastikan status draft
    if ref.Status != "draft" {
        return errors.New("only draft achievements can be deleted")
    }

    // 3. Soft delete MongoDB
    err = s.Repo.SoftDeleteMongo(ctx, ref.MongoAchievementID)
    if err != nil {
        return err
    }

    // 4. Soft delete PostgreSQL
    return s.Repo.SoftDelete(ctx, id, userID)
}

func (s *AchievementService) Reject(ctx context.Context, id string, note string) error {

    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    if ref.Status != "submitted" {
        return errors.New("only submitted achievements can be rejected")
    }

    // Pakai lecturerID untuk validasi advisor
    lecturerID := ctx.Value("lecturerID").(string)

    allowed, err := s.Repo.IsAdvisorOf(ctx, lecturerID, ref.StudentID)
    if err != nil {
        return err
    }
    if !allowed {
        return errors.New("forbidden: not advisor of this student")
    }

    // Pakai userID untuk FK verified_by
    userID := ctx.Value("userID").(string)

    return s.Repo.RejectAchievement(ctx, id, userID, note)
}

func (s *AchievementService) ListForStudent(ctx context.Context, studentID string) ([]models.Achievement, error) {
    return s.Repo.ListByStudent(ctx, studentID)
}

func (s *AchievementService) ListForAdvisor(ctx context.Context, advisorID string) ([]models.AchievementResponse, error) {

    refs, err := s.Repo.ListByAdvisorStudents(ctx, advisorID)
    if err != nil {
        return nil, err
    }

    responses := []models.AchievementResponse{}

    for _, r := range refs {

        // Ambil data detail di Mongo
        detail, err := s.Repo.GetMongoByID(ctx, r.MongoAchievementID)
        if err != nil {
            continue // skip error
        }

        responses = append(responses, models.AchievementResponse{
            ID:          r.ID,
            MongoID:     r.MongoAchievementID,
            StudentID:   r.StudentID,
            Title:       detail.Title,
            Description: detail.Description,
            Category:    detail.AchievementType,
            Status:      r.Status,
            CreatedAt:   r.CreatedAt.Format(time.RFC3339),
        })
    }

    return responses, nil
}


func (s *AchievementService) GetHistory(ctx context.Context, id string) (*models.AchievementHistoryResponse, error) {

    // Ambil data achievement dari PostgreSQL
    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    role := ctx.Value("role").(string)
    // userID := ctx.Value("userID").(string)
    studentID := ctx.Value("studentID").(string)
    lecturerID := ctx.Value("lecturerID").(string)

    switch role {

    // Mahasiswa → hanya bisa lihat history miliknya
    case "Mahasiswa":
        if ref.StudentID != studentID {
            return nil, errors.New("forbidden")
        }


    // Dosen Wali → hanya mahasiswa bimbingannya
    case "Dosen Wali":
        allowed, err := s.Repo.IsAdvisorOf(ctx, lecturerID, ref.StudentID)
        if err != nil {
            return nil, err
        }
        if !allowed {
            return nil, errors.New("forbidden")
        }

    // Admin → selalu boleh
    case "Admin":
        // no validation needed
    // Role lain (jika ada)
    default:
        return nil, errors.New("forbidden")
    }
    // BANGUN HISTORY DARI DATABASE
    history := []models.AchievementHistoryItem{}

    // DRAFT
    history = append(history, models.AchievementHistoryItem{
        Status:    "draft",
        ChangedAt: ref.CreatedAt,
        ChangedBy: ref.StudentID,
    })

    // SUBMITTED
    if ref.SubmittedAt != nil {
        history = append(history, models.AchievementHistoryItem{
            Status:    "submitted",
            ChangedAt: *ref.SubmittedAt,
            ChangedBy: ref.StudentID,
        })
    }

    // VERIFIED
    if ref.VerifiedAt != nil && ref.VerifiedBy != nil && ref.RejectionNote == nil {
        history = append(history, models.AchievementHistoryItem{
            Status:    "verified",
            ChangedAt: *ref.VerifiedAt,
            ChangedBy: *ref.VerifiedBy,
        })
    }

    // REJECTED
    if ref.VerifiedAt != nil && ref.VerifiedBy != nil && ref.RejectionNote != nil {
        history = append(history, models.AchievementHistoryItem{
            Status:    "rejected",
            ChangedAt: *ref.VerifiedAt,
            ChangedBy: *ref.VerifiedBy,
            Note:      *ref.RejectionNote,
        })
    }

    // DELETED
    if ref.DeletedAt != nil && ref.DeletedBy != nil {
        history = append(history, models.AchievementHistoryItem{
            Status:    "deleted",
            ChangedAt: *ref.DeletedAt,
            ChangedBy: *ref.DeletedBy,
        })
    }

    return &models.AchievementHistoryResponse{
        AchievementID: ref.ID,
        History:       history,
    }, nil
}

func (s *AchievementService) UploadAttachment(ctx context.Context, id string, att models.AttachmentMongo) error {

    ref, err := s.Repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    if ref.Status == "verified" || ref.Status == "deleted" {
        return errors.New("cannot upload attachment to verified or deleted achievement")
    }

    return s.Repo.AddAttachment(ctx, ref.MongoAchievementID, att)
}

