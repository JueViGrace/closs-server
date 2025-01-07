package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type AuthStore interface {
	SignIn(r *types.SignInRequest) (*types.AuthResponse, error)
	SignUp(r *types.SignUpRequest) (*types.AuthResponse, error)
	Refresh(r *types.RefreshRequest) (*types.AuthResponse, error)
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

func (s *authStore) SignIn(r *types.SignInRequest) (string, error) {
	// user, err := s.db.GetUserByUsername(s.ctx, r.Username)

	return "", nil
}

func (s *authStore) SignUp(r *types.SignUpRequest) (string, error) {
	return "", nil
}

func (s *authStore) RecoverPassword(r *types.RecoverPasswordResquest) (string, error) {
	return "", nil
}
