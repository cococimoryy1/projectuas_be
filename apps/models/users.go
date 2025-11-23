package models

type User struct {
    ID          string `json:"id"`
    Username    string `json:"username"`
    Email       string `json:"email"`
    PasswordHash string `json:"password_hash"`
    RoleID      string `json:"role_id"`
    RoleName    string `json:"role_name"`
    IsActive    bool   `json:"is_active"` // Tambah dari SRS hal.4
}

type Student struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    StudentID   string `json:"student_id"`
    Program     string `json:"program"`
    AcademicYear string `json:"academic_year"`
    AdvisorID   string `json:"advisor_id"`
}