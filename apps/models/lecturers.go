package models

type LecturerListResponse struct {
    ID           string `json:"id"`
    LecturerID   string `json:"lecturer_id"`
    FullName     string `json:"full_name"`
    Email        string `json:"email"`
    Department   string `json:"department"`
    AdviseeCount int    `json:"advisee_count"`
}

type LecturerDetailResponse struct {
    ID         string `json:"id"`
    LecturerID string `json:"lecturer_id"`
    FullName   string `json:"full_name"`
    Email      string `json:"email"`
    Department string `json:"department"`
}
type AdviseeResponse struct {
    StudentID     string `json:"student_id"`
    FullName      string `json:"full_name"`
    ProgramStudy  string `json:"program_study"`
    AcademicYear  string `json:"academic_year"`
}
