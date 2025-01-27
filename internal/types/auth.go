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
	Role     Role
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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	Token string `json:"refreshToken" validate:"required"`
}

type RecoverPasswordResquest struct {
	Username string `json:"username" validate:"required"`
}
