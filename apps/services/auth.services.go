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
    UserRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
    return &AuthService{UserRepo: repo}
}


func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
    if req.Username == "" || req.Password == "" {
        return nil, errors.New("username and password required")
    }

    user, err := s.UserRepo.FindByUsernameOrEmail(ctx, req.Username)
    if err != nil || !user.IsActive {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    perms, err := s.UserRepo.GetPermissionsByRoleID(ctx, user.RoleID)
    if err != nil {
        return nil, errors.New("permission load error")
    }

    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return nil, errors.New("jwt secret missing")
    }

    // === access token ===
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

    // === refresh token ===
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
    })
    // SIGN refresh token dengan JWT_SECRET juga
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

// ======================
// REFRESH TOKEN
// ======================
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

    user, err := s.UserRepo.FindByID(ctx, userID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    perms, err := s.UserRepo.GetPermissionsByRoleID(ctx, user.RoleID)
    if err != nil {
        return nil, errors.New("permission load error")
    }

    // === generate ACCESS TOKEN baru ===
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
        RefreshToken: refreshToken, // tetap pakai refresh lama
        User: models.UserResponse{
            ID:          user.ID,
            Username:    user.Username,
            Email:       user.Email,
            Role:        user.RoleName,
            Permissions: perms,
        },
    }, nil
}


// Logout: Simple success (optional blacklist token)
func (s *AuthService) Logout(ctx context.Context) error {
    // Optional: Blacklist token di Redis/DB (implement nanti)
    return nil
}

// Profile: Get user profile from JWT (no DB, fast)
func (s *AuthService) Profile(ctx context.Context, userID string) (*models.UserResponse, error) {
    // Load full user dari repo (optional, jika JWT cukup)
    user, err := s.UserRepo.FindByID(ctx, userID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    perms, err := s.UserRepo.GetPermissionsByRoleID(ctx, user.RoleID)
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