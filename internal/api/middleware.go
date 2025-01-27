package api

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.Role == types.RoleAdmin {
		res := types.RespondForbbiden("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	_, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authenticatedAdminHandler(handler types.AuthHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondUnauthorized(nil, "you are not authorized to access this endpoint")
			return c.Status(res.Status).JSON(res)
		}

		if data.Role != types.RoleAdmin {
			res := types.RespondForbbiden(nil, "forbbiden resource")
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func (a *api) authenticatedHandler(handler types.AuthHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondUnauthorized(nil, "you are not authorized to access this endpoint")
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*types.AuthData, error) {
	jwt, err := extractJWTFromHeader(c, func(u uuid.UUID) {
		db.SessionStore().DeleteSessionById(u)
	})
	if err != nil {
		return nil, err
	}

	session, err := db.SessionStore().GetSessionById(jwt.Claims.UserId)
	if err != nil {
		return nil, err
	}

	user, err := db.UserStore().GetUserById(session.UserId)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	return &types.AuthData{
		UserId:   userId,
		Role:     user.Role,
		Username: user.Username,
		Code:     user.Code,
	}, nil
}

func extractJWTFromHeader(c *fiber.Ctx, expired func(uuid.UUID)) (*types.JwtData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := util.ValidateJWT(tokenString)
	if err != nil {
		return nil, errors.New("permission denied")
	}

	if !token.Valid {
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(util.JWTClaims)
	if !ok {
		return nil, errors.New("permission denied")
	}

	if claims.ExpiresAt.Time.UTC().Unix() < time.Now().UTC().Unix() {
		expired(claims.UserId)
		return nil, errors.New("permision denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != util.Issuer {
		return nil, errors.New("permision denied")
	}

	return &types.JwtData{
		Token:  token,
		Claims: claims,
	}, nil
}
