package models

type User struct {
    ID                  string  `json:"id"`
    Username            string  `json:"username"`
    Email               string  `json:"email"`
    PasswordHash        string  `json:"password_hash"`
    FullName            string `json:"full_name"`
    RoleID              string  `json:"role_id"`
    RoleName            string  `json:"role_name"`
    IsActive            bool    `json:"is_active"` // Tambah dari SRS hal.4
}

type UserResponse struct {
    ID                  string   `json:"id"`
    Username            string   `json:"username"`
    Email               string   `json:"email"`
    Role                string   `json:"role"`
    Permissions         []string `json:"permissions"`
}

type CreateUserRequest struct {
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    FullName  string `json:"full_name"`
    RoleID    string `json:"role_id"`
}

type UpdateUserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    FullName string `json:"full_name"`
}

type UpdateRoleRequest struct {
    RoleID              string `json:"role_id"`
}

type EmptyRequest struct{}
