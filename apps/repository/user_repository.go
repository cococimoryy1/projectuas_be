package repository

import (
    "BE_PROJECTUAS/database"
    "BE_PROJECTUAS/apps/models"
    "context"
)

type userRepo struct{}

func NewUserRepository() UserRepository {
    return &userRepo{}
}

func (r *userRepo) FindByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error) {
    query := `
        SELECT u.id, u.username, u.email, u.password_hash,
               u.role_id, r.name AS role_name, u.is_active
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.username = $1 OR LOWER(u.email) = LOWER($1)
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

func (r *userRepo) GetPermissionsByRoleID(ctx context.Context, roleID string) ([]string, error) {
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
        if err := rows.Scan(&name); err != nil { // Fix: assign err
            return nil, err
        }
        list = append(list, name)
    }

    if err := rows.Err(); err != nil { // Tambah: Handle iteration error
        return nil, err
    }

    return list, nil
}
func (r *userRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
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
