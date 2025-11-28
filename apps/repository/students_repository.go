package repository

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"
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
    LEFT JOIN users u2 ON u2.id = l.user_id  -- ambil nama dosen
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
        if err := rows.Scan(
            &s.ID,
            &s.StudentID,
            &s.Username,
            &s.Email,
            &s.ProgramStudy,
            &s.AcademicYear,
            &s.AdvisorID,
            &s.AdvisorName,
        ); err != nil {
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
        SELECT COUNT(*) FROM students 
        WHERE id = $1 AND advisor_id = $2
    `

    var count int
    err := database.PostgresDB.QueryRowContext(ctx, query, studentID, advisorID).Scan(&count)

    return count > 0, err
}
