package data

import (
	"context"
	"errors"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (*types.AuthResponse, error)
	Refresh(r *types.RefreshRequest, d *types.AuthData) (*types.AuthResponse, error)
	RecoverPassword(r *types.RecoverPasswordResquest) (*types.AuthResponse, error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx, s.queries)
}

type authStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewAuthStore(ctx context.Context, db *db.Queries) AuthStore {
	return &authStore{
		ctx: ctx,
		db:  db,
	}
}

// todo: refactor functions
// todo: search for existing session and delete it?
// if not add support for limited number of sessions
func (s *authStore) SignIn(r *types.SignInRequest) (*types.AuthResponse, error) {
	user, err := s.db.GetUserByUsername(s.ctx, r.Username)
	if err != nil {
		return nil, err
	}

	if user.DeletedAt.Valid {
		return nil, errors.New("this user was deleted")
	}

	if !util.ValidatePassword(r.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	res, err := createTokens(&user)
	if err != nil {
		return nil, err
	}

	err = s.db.CreateSession(s.ctx, db.CreateSessionParams{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		Username:     user.Username,
		UserID:       user.ID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authStore) Refresh(r *types.RefreshRequest, d *types.AuthData) (*types.AuthResponse, error) {
	user, err := s.db.GetUserById(s.ctx, d.UserId.String())
	if err != nil {
		return nil, err
	}

	res, err := createTokens(&user)
	if err != nil {
		return nil, err
	}

	err = s.db.UpdateSession(s.ctx, db.UpdateSessionParams{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		UserID:       user.ID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authStore) RecoverPassword(r *types.RecoverPasswordResquest) (*types.AuthResponse, error) {
	return nil, nil
}

func createTokens(user *db.ClossUser) (*types.AuthResponse, error) {
	accessToken, err := util.CreateAccessToken(user.ID, user.Username, user.Codigo)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.CreateRefreshToken(user.ID, user.Username, user.Codigo)
	if err != nil {
		return nil, err
	}

	return &types.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
