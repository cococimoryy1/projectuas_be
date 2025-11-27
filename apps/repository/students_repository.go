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
            COALESCE(l.lecturer_id, '') AS advisor_name
        FROM students s
        JOIN users u ON u.id = s.user_id
        LEFT JOIN lecturers l ON l.id = s.advisor_id
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
