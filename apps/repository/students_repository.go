package repository

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"
    "time"
)



type studentRepo struct {}

func NewStudentRepository() StudentRepository {
    return &studentRepo{}
}

func (r *studentRepo) ListStudents(ctx context.Context) ([]models.StudentListResponse, error) {

    query := `
        SELECT 
            s.id,
            s.student_id,
            u.username,
            u.email,
            s.program_study,
            s.academic_year,
            s.advisor_id,
            COALESCE(u2.full_name, '') AS advisor_name
        FROM students s
        JOIN users u ON u.id = s.user_id
        LEFT JOIN lecturers l ON l.id = s.advisor_id
        LEFT JOIN users u2 ON u2.id = l.user_id
        ORDER BY s.student_id;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.StudentListResponse

    for rows.Next() {
        var s models.StudentListResponse
        err := rows.Scan(
            &s.ID,
            &s.StudentID,
            &s.Username,
            &s.Email,
            &s.ProgramStudy,
            &s.AcademicYear,
            &s.AdvisorID,
            &s.AdvisorName,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, s)
    }

    return list, nil
}

func (r *studentRepo) GetByID(ctx context.Context, id string) (*models.StudentDetailResponse, error) {

    query := `
        SELECT 
            s.id,
            s.student_id,
            u.full_name,
            u.email,
            s.program_study,
            s.academic_year,
            s.advisor_id,
            COALESCE(u2.full_name, '') AS advisor_name
        FROM students s
        JOIN users u ON u.id = s.user_id
        LEFT JOIN lecturers l ON l.id = s.advisor_id
        LEFT JOIN users u2 ON u2.id = l.user_id
        WHERE s.id = $1
        LIMIT 1;
    `

    var data models.StudentDetailResponse

    err := database.PostgresDB.QueryRowContext(ctx, query, id).Scan(
        &data.ID,
        &data.StudentID,
        &data.FullName,
        &data.Email,
        &data.ProgramStudy,
        &data.AcademicYear,
        &data.AdvisorID,
        &data.AdvisorName,
    )

    if err != nil {
        return nil, err
    }

    return &data, nil
}

func (r *studentRepo) IsAdvisorOf(ctx context.Context, advisorID string, studentID string) (bool, error) {

    query := `
        SELECT COUNT(*)
        FROM students 
        WHERE id = $1 AND advisor_id = $2
    `

    var count int
    err := database.PostgresDB.QueryRowContext(ctx, query, studentID, advisorID).Scan(&count)
    return count > 0, err
}
func (r *achievementRepo) ListByStudentID(ctx context.Context, studentID string) ([]models.AchievementResponse, error) {

    query := `
        SELECT id, mongo_achievement_id, status, created_at
        FROM achievement_references
        WHERE student_id = $1;
    `
    rows, err := database.PostgresDB.QueryContext(ctx, query, studentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []models.AchievementResponse

    for rows.Next() {

        var id, mongoID, status string
        var created time.Time

        if err := rows.Scan(&id, &mongoID, &status, &created); err != nil {
            return nil, err
        }

        detail, _ := r.GetMongoByID(ctx, mongoID)

        list = append(list, models.AchievementResponse{
            ID:          id,
            MongoID:     mongoID,
            StudentID:   studentID,
            Title:       detail.Title,
            Description: detail.Description,
            Category:    detail.AchievementType,
            Status:      status,
            CreatedAt:   created.Format(time.RFC3339),
        })
    }

    return list, nil
}
