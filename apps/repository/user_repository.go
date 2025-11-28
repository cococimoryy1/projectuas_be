package repository

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"
    "context"

    "github.com/google/uuid"
)


type userRepo struct{}

func NewUserRepository() UserRepository {
    return &userRepo{}
}


// GET ALL USERS
func (r *userRepo) ListUsers(ctx context.Context) ([]models.User, error) {
    query := `
        SELECT u.id, u.username, u.email, u.password_hash,
               u.role_id, r.name AS role_name, u.is_active
        FROM users u
        JOIN roles r ON u.role_id = r.id
        ORDER BY u.created_at DESC;
    `

    rows, err := database.PostgresDB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var u models.User
        if err := rows.Scan(
            &u.ID, &u.Username, &u.Email, &u.PasswordHash,
            &u.RoleID, &u.RoleName, &u.IsActive,
        ); err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, nil
}

// GET BY ID
func (r *userRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    query := `
        SELECT u.id, u.username, u.email, u.password_hash,
               u.role_id, r.name AS role_name, u.is_active
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.id = $1
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

// CREATE USER
func (r *userRepo) CreateUser(ctx context.Context, req models.CreateUserRequest) (string, error) {
    query := `
    INSERT INTO users (id, username, email, password_hash, full_name, role_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
    RETURNING id;
    `

    // generate UUID
    newID := uuid.NewString()

    var id string
    err := database.PostgresDB.QueryRowContext(ctx, query,
        newID,         // $1 → id
        req.Username,  // $2
        req.Email,     // $3
        req.Password,  // $4
        req.FullName,  // $5
        req.RoleID,    // $6
    ).Scan(&id)

    if err != nil {
        return "", err
    }

    return id, nil
}

// UPDATE
func (r *userRepo) UpdateUser(ctx context.Context, id string, req models.UpdateUserRequest) error {
    query := `
        UPDATE users
        SET username=$1, email=$2, full_name=$3, updated_at=NOW()
        WHERE id=$4
    `
    _, err := database.PostgresDB.ExecContext(ctx, query,
        req.Username,
        req.Email,
        req.FullName,
        id,
    )
    return err
}

// DELETE
func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
    query := `DELETE FROM users WHERE id=$1;`
    _, err := database.PostgresDB.ExecContext(ctx, query, id)
    return err
}

// UPDATE ROLE ONLY (PUT /users/:id/role)

func (r *userRepo) UpdateUserRole(ctx context.Context, id string, roleID string) error {
    query := `UPDATE users SET role_id=$1 WHERE id=$2;`
    _, err := database.PostgresDB.ExecContext(ctx, query, roleID, id)
    return err
}
