package repository

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"
)

type lecturerRepo struct{}

func NewLecturerRepository() LecturerRepository {
    return &lecturerRepo{}
}

func (r *lecturerRepo) ListLecturers(ctx context.Context) ([]models.LecturerListResponse, error) {

    query := `
        SELECT 
            l.id,
            l.lecturer_id,
            u.full_name,
            u.email,
            l.department,
            COALESCE((
                SELECT COUNT(*) FROM students s WHERE s.advisor_id = l.id
            ), 0) AS advisee_count
        FROM lecturers l
        JOIN users u ON u.id = l.user_id
        ORDER BY u.full_name;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    list := []models.LecturerListResponse{}

    for rows.Next() {
        var item models.LecturerListResponse
        if err := rows.Scan(
            &item.ID,
            &item.LecturerID,
            &item.FullName,
            &item.Email,
            &item.Department,
            &item.AdviseeCount,
        ); err != nil {
            return nil, err
        }
        list = append(list, item)
    }

    return list, nil
}
func (r *lecturerRepo) GetByID(ctx context.Context, id string) (*models.LecturerDetailResponse, error) {
    query := `
        SELECT l.id, l.lecturer_id, u.full_name, u.email, l.department
        FROM lecturers l
        JOIN users u ON u.id = l.user_id
        WHERE l.id = $1
        LIMIT 1;
    `

    var data models.LecturerDetailResponse

    err := database.PostgresDB.QueryRowContext(ctx, query, id).Scan(
        &data.ID,
        &data.LecturerID,
        &data.FullName,
        &data.Email,
        &data.Department,
    )

    if err != nil {
        return nil, err
    }

    return &data, nil
}

func (r *lecturerRepo) ListAdvisees(ctx context.Context, lecturerID string) ([]models.StudentListResponse, error) {

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
        LEFT JOIN users u2 ON u2.id = s.advisor_id
        WHERE s.advisor_id = $1;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query, lecturerID)
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
            &s.Username,      // u.full_name
            &s.Email,
            &s.ProgramStudy,
            &s.AcademicYear,
            &s.AdvisorID,
            &s.AdvisorName,   // sudah aman (COALESCE)
        ); err != nil {
            return nil, err
        }

        list = append(list, s)
    }

    return list, nil
}



func (r *lecturerRepo) GetAdvisorByStudentID(ctx context.Context, studentID string) (string, error) {
    query := `SELECT advisor_id FROM students WHERE id = $1 LIMIT 1`

    var advisorID string
    err := database.PostgresDB.QueryRowContext(ctx, query, studentID).Scan(&advisorID)
    return advisorID, err
}
