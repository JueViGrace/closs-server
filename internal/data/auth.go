package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/types"
)

type AuthStore interface {
	SignIn(r types.SignInRequest) (string, error)
	SignUp(r types.SignUpRequest) (string, error)
	RecoverPassword(username string) (string, error)
}

func (s *storage) AuthStore() AuthStore {
	return NewAuthStore(s.ctx)
}

type authStore struct {
	ctx context.Context
}

func NewAuthStore(ctx context.Context) AuthStore {
	return &authStore{
		ctx: ctx,
	}
}

func (s *authStore) SignIn(r types.SignInRequest) (string, error) {
	return "", nil
}

func (s *authStore) SignUp(r types.SignUpRequest) (string, error) {
	return "", nil
}

func (s *authStore) RecoverPassword(username string) (string, error) {
	return "", nil
}
