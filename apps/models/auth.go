package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
    UserID              string      `json:"user_id"`
    StudentID           string      `json:"student_id"` 
    LecturerID          string      `json:"lecturerId"`
    Username            string      `json:"username"`
    RoleName            string      `json:"role"`
    Permissions         []string    `json:"permissions"`
    jwt.RegisteredClaims
}

type LoginRequest struct {
    Username            string      `json:"username"`
    Password            string      `json:"password"`
}

type LoginResponse struct {
    Token               string       `json:"token"`
    RefreshToken        string       `json:"refreshToken"` 
    User                UserResponse `json:"user"`
}

type RefreshRequest struct {
    RefreshToken string `json:"refreshToken"`
}
