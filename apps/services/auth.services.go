package services

import (
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/apps/repository"
    "context"
    "errors"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    AuthRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
    return &AuthService{AuthRepo: repo}
}

// ====================
// LOGIN
// ====================
func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
    if req.Username == "" || req.Password == "" {
        return nil, errors.New("username and password required")
    }

    user, err := s.AuthRepo.FindByUsernameOrEmail(ctx, req.Username)
    if err != nil || !user.IsActive {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    perms, err := s.AuthRepo.GetPermissionsByRoleID(ctx, user.RoleID)
    if err != nil {
        return nil, errors.New("permission load error")
    }

    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return nil, errors.New("jwt secret missing")
    }

    // === Access Token ===
    claims := models.JwtCustomClaims{
        UserID:      user.ID,
        Username:    user.Username,
        RoleName:    user.RoleName,
        Permissions: perms,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(secret))
    if err != nil {
        return nil, errors.New("token generation failed")
    }

    // === Refresh Token ===
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
    })
    refreshStr, _ := refreshToken.SignedString([]byte(secret))

    return &models.LoginResponse{
        Token:        signedToken,
        RefreshToken: refreshStr,
        User: models.UserResponse{
            ID:          user.ID,
            Username:    user.Username,
            Email:       user.Email,
            Role:        user.RoleName,
            Permissions: perms,
        },
    }, nil
}

// ====================
// REFRESH
// ====================
func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (*models.LoginResponse, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return nil, errors.New("jwt secret missing")
    }

    token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("invalid refresh token")
    }

    claims := token.Claims.(jwt.MapClaims)
    userID := claims["sub"].(string)

    user, err := s.AuthRepo.FindByID(ctx, userID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    perms, err := s.AuthRepo.GetPermissionsByRoleID(ctx, user.RoleID)
    if err != nil {
        return nil, errors.New("permission load error")
    }

    accessClaims := models.JwtCustomClaims{
        UserID:      user.ID,
        Username:    user.Username,
        RoleName:    user.RoleName,
        Permissions: perms,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    signedAccess, err := accessToken.SignedString([]byte(secret))
    if err != nil {
        return nil, errors.New("token generation failed")
    }

    return &models.LoginResponse{
        Token:        signedAccess,
        RefreshToken: refreshToken,
        User: models.UserResponse{
            ID:          user.ID,
            Username:    user.Username,
            Email:       user.Email,
            Role:        user.RoleName,
            Permissions: perms,
        },
    }, nil
}

func (s *AuthService) Logout(ctx context.Context) error {
    return nil // implement blacklist if needed
}

func (s *AuthService) Profile(ctx context.Context, userID string) (*models.UserResponse, error) {
    user, err := s.AuthRepo.FindByID(ctx, userID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    perms, err := s.AuthRepo.GetPermissionsByRoleID(ctx, user.RoleID)
    if err != nil {
        return nil, errors.New("permission load error")
    }

    return &models.UserResponse{
        ID:          user.ID,
        Username:    user.Username,
        Email:       user.Email,
        Role:        user.RoleName,
        Permissions: perms,
    }, nil
}
