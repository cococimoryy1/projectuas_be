package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
    UserID      string   `json:"user_id"`
    StudentID   string   `json:"student_id"` // ← Wajib
    Username    string   `json:"username"`
    RoleName    string   `json:"role"`
    Permissions []string `json:"permissions"`
    jwt.RegisteredClaims
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token        string       `json:"token"`
    RefreshToken string       `json:"refreshToken"` // Optional, bisa tambah nanti
    User         UserResponse `json:"user"`
}


