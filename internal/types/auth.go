package types

import (
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthHandler = func(c *fiber.Ctx, d *AuthData) error

type AuthData struct {
	UserId   uuid.UUID
	Role     string
	Username string
	Code     string
}

type JwtData struct {
	Token  *jwt.Token
	Claims util.JWTClaims
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	Token string `json:"refreshToken"`
}

type RecoverPasswordResquest struct {
	Username string `json:"username"`
}
