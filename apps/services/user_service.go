package services

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    "context"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    Repo repository.UserRepository
}
func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) List(ctx context.Context) (*[]models.User, error) {
    users, err := s.Repo.ListUsers(ctx)
    if err != nil {
        return nil, err
    }
    return &users, nil
}

func (s *UserService) Get(ctx context.Context, id string) (*models.User, error) {
    return s.Repo.GetUserByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) (*string, error) {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
    req.Password = string(hashed)
    id, err := s.Repo.CreateUser(ctx, req)
    if err != nil {
        return nil, err
    }
    return &id, nil
}

func (s *UserService) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*string, error) {
    err := s.Repo.UpdateUser(ctx, id, req)
    if err != nil {
        return nil, err
    }
    msg := "updated"
    return &msg, nil
}

func (s *UserService) Delete(ctx context.Context, id string) (*string, error) {
    err := s.Repo.DeleteUser(ctx, id)
    if err != nil {
        return nil, err
    }
    msg := "deleted"
    return &msg, nil
}

func (s *UserService) UpdateRole(ctx context.Context, id string, req models.UpdateRoleRequest) (*string, error) {
    err := s.Repo.UpdateUserRole(ctx, id, req.RoleID)
    if err != nil {
        return nil, err
    }

    msg := "role updated"
    return &msg, nil
}
