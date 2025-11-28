package models

type StudentListResponse struct {
    ID              string `json:"id"`
    StudentID       string `json:"student_id"`
    Username        string `json:"username"`
    Email           string `json:"email"`
    ProgramStudy    string `json:"program_study"`
    AcademicYear    string `json:"academic_year"`
    AdvisorID       string `json:"advisor_id"`
    AdvisorName     string `json:"advisor_name"`
}

type StudentDetailResponse struct {
    ID              string `json:"id"`
    StudentID       string `json:"student_id"`
    FullName        string `json:"full_name"`
    Email           string `json:"email"`
    ProgramStudy    string `json:"program_study"`
    AcademicYear    string `json:"academic_year"`
    AdvisorID       string `json:"advisor_id"`
    AdvisorName     string `json:"advisor_name"`
}

type Student struct {
    ID            string `json:"id"`
    UserID        string `json:"user_id"`
    StudentID     string `json:"student_id"`
    FullName      string `json:"full_name"`
    ProgramStudy  string `json:"program_study"`
    AcademicYear  string `json:"academic_year"`
    AdvisorID     string `json:"advisor_id"`
}
