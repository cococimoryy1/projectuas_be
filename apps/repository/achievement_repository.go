package repository

import (
    "BE_PROJECTUAS/database"
    "BE_PROJECTUAS/apps/models"
    "context"

    "github.com/google/uuid"
)

type achievementRepo struct{}

func NewAchievementRepository() AchievementRepository {
    return &achievementRepo{}
}

func (r *achievementRepo) CreateAchievementReference(ctx context.Context, a models.Achievement) (string, error) {
    // Auto-generate ID jika kosong (SRS hal.4 UUID)
    if a.ID == "" {
        a.ID = uuid.New().String()
    }

    query := `
        INSERT INTO achievement_references 
        (id, student_id, mongo_achievement_id, status)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

    var id string
    err := database.PostgresDB.QueryRowContext(ctx, query, a.ID, a.StudentID, a.MongoID, a.Status).Scan(&id)
    return id, err
}

func (r *achievementRepo) UpdateStatus(ctx context.Context, id string, status string) error {
    query := `UPDATE achievement_references SET status = $1 WHERE id = $2;`
    _, err := database.PostgresDB.ExecContext(ctx, query, status, id)
    return err
}

func (r *achievementRepo) ListByStudent(ctx context.Context, studentID string) ([]models.Achievement, error) {
    query := `
        SELECT id, student_id, mongo_achievement_id, status
        FROM achievement_references
        WHERE student_id = $1;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.Achievement
    for rows.Next() {
        var a models.Achievement
        if err := rows.Scan(&a.ID, &a.StudentID, &a.MongoID, &a.Status); err != nil { // Fix: assign err
            return nil, err
        }
        list = append(list, a)
    }

    if err := rows.Err(); err != nil { // Tambah
        return nil, err
    }

    return list, nil
}

func (r *achievementRepo) ListByAdvisorStudents(ctx context.Context, advisorID string) ([]models.Achievement, error) {
    query := `
        SELECT ar.id, ar.student_id, ar.mongo_achievement_id, ar.status
        FROM achievement_references ar
        JOIN students s ON s.id = ar.student_id
        WHERE s.advisor_id = $1;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query, advisorID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.Achievement
    for rows.Next() {
        var a models.Achievement
        if err := rows.Scan(&a.ID, &a.StudentID, &a.MongoID, &a.Status); err != nil {
            return nil, err
        }
        list = append(list, a)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return list, nil
}