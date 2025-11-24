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

type UserResponse struct {
    ID          string   `json:"id"`
    Username    string   `json:"username"`
    Email       string   `json:"email"`
    Role        string   `json:"role"`
    Permissions []string `json:"permissions"`
}

type Student struct {
    ID          string `json:"id"`
    UserID      string `json:"user_id"`
    StudentID   string `json:"student_id"`
    Program     string `json:"program"`
    AcademicYear string `json:"academic_year"`
    AdvisorID   string `json:"advisor_id"`
}
type CreateUserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    RoleID   string `json:"role_id"`
}

type UpdateUserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    RoleID   string `json:"role_id"`
}

type UpdateRoleRequest struct {
    RoleID string `json:"role_id"`
}

type EmptyRequest struct{}
