package utils

import (
    "BE_PROJECTUAS/apps/models"
    "os"
    

    "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims models.JwtCustomClaims) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "defaultsecret" // fallback dev mode
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*models.JwtCustomClaims, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "defaultsecret"
    }

    token, err := jwt.ParseWithClaims(
        tokenString,
        &models.JwtCustomClaims{},
        func(t *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        },
    )
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*models.JwtCustomClaims)
    if !ok || !token.Valid {
        return nil, jwt.ErrInvalidKey
    }

    return claims, nil
}
