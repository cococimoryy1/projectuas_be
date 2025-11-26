package repository

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"
    "context"
)

type authRepo struct{}

func NewAuthRepository() AuthRepository {
    return &authRepo{}
}

func (r *authRepo) FindByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error) {
    query := `
        SELECT u.id, u.username, u.email, u.password_hash,
               u.role_id, r.name AS role_name, u.is_active
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE (u.username = $1 OR LOWER(u.email) = LOWER($1))
          AND u.is_active = true
        LIMIT 1;
    `

    var u models.User
    err := database.PostgresDB.QueryRowContext(ctx, query, identifier).Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.RoleID, &u.RoleName, &u.IsActive,
    )
    if err != nil {
        return nil, err
    }

    return &u, nil
}
func (r *authRepo) GetStudentByUserID(ctx context.Context, userID string) (string, error) {
    query := `SELECT id FROM students WHERE user_id = $1 LIMIT 1`

    var studentID string
    err := database.PostgresDB.QueryRowContext(ctx, query, userID).Scan(&studentID)
    if err != nil {
        return "", err
    }
    return studentID, nil
}


func (r *authRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
    query := `
        SELECT u.id, u.username, u.email, u.password_hash,
               u.role_id, r.name AS role_name, u.is_active
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.id = $1
        LIMIT 1;
    `

    var u models.User
    err := database.PostgresDB.QueryRowContext(ctx, query, id).Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash,
        &u.RoleID, &u.RoleName, &u.IsActive,
    )
    if err != nil {
        return nil, err
    }

    return &u, nil
}

func (r *authRepo) GetPermissionsByRoleID(ctx context.Context, roleID string) ([]string, error) {
    query := `
        SELECT p.name
        FROM permissions p
        JOIN role_permissions rp ON p.id = rp.permission_id
        WHERE rp.role_id = $1;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query, roleID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []string

    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        list = append(list, name)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return list, nil
}
func (r *authRepo) GetLecturerByUserID(ctx context.Context, userID string) (string, error) {
    query := `SELECT id FROM lecturers WHERE user_id = $1 LIMIT 1`

    var lecturerID string
    err := database.PostgresDB.QueryRowContext(ctx, query, userID).Scan(&lecturerID)
    if err != nil {
        return "", err
    }
    return lecturerID, nil
}
