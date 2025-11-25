package repository

import (
    "BE_PROJECTUAS/database"
    "BE_PROJECTUAS/apps/models"
    "context"

    // "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type achievementRepo struct{}

func NewAchievementRepository() AchievementRepository {
    return &achievementRepo{}
}

func (r *achievementRepo) CreateAchievementReference(ctx context.Context, a models.Achievement) (string, error) {
    query := `
        INSERT INTO achievement_references 
        (id, student_id, mongo_achievement_id, status, created_at, updated_at)
        VALUES ($1,$2,$3,'draft',NOW(),NOW())
        RETURNING id;
    `
    var id string
    err := database.PostgresDB.QueryRowContext(ctx, query,
        a.ID,
        a.StudentID,
        a.MongoAchievementID,
    ).Scan(&id)

    if err != nil {
        return "", err
    }

    return id, nil
}

func (r *achievementRepo) InsertMongoAchievement(ctx context.Context, doc models.AchievementMongo) (string, error) {
    col := database.MongoDB.Collection("achievements")

    result, err := col.InsertOne(ctx, doc)
    if err != nil {
        return "", err
    }

    oid := result.InsertedID.(primitive.ObjectID)
    return oid.Hex(), nil
}

func (r *achievementRepo) GetByID(ctx context.Context, id string) (*models.Achievement, error) {
    query := `
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by, rejection_note
        FROM achievement_references
        WHERE id = $1;
    `
    row := database.PostgresDB.QueryRowContext(ctx, query, id)

    var a models.Achievement
    err := row.Scan(
        &a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
        &a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy, &a.RejectionNote,
    )
    if err != nil {
        return nil, err
    }
    return &a, nil
}
func (r *achievementRepo) VerifyAchievement(ctx context.Context, id, lecturerID string) error {
    query := `
        UPDATE achievement_references
        SET status = 'verified',
            verified_at = NOW(),
            verified_by = $2
        WHERE id = $1;
    `
    _, err := database.PostgresDB.ExecContext(ctx, query, id, lecturerID)
    return err
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
        if err := rows.Scan(&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status); err != nil { // Fix: assign err
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
        if err := rows.Scan(&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status); err != nil {
            return nil, err
        }
        list = append(list, a)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return list, nil
}
func (r *achievementRepo) IsAdvisorOf(ctx context.Context, lecturerID string, studentID string) (bool, error) {
    query := `
        SELECT COUNT(*) 
        FROM students 
        WHERE id = $1 AND advisor_id = $2
    `
    var count int
    err := database.PostgresDB.QueryRowContext(ctx, query, studentID, lecturerID).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
