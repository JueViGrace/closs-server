package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() ([]*types.User, error)
	GetUser(id uuid.UUID) (*types.User, error)
	UpdateUser(r *types.UpdateUserRequest) (string, error)
	DeleteUser(id uuid.UUID) error
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.queries)
}

type userStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewUserStore(ctx context.Context, db *db.Queries) UserStore {
	return &userStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStore) GetUsers() ([]*types.User, error) {
	return nil, nil
}

func (s *userStore) GetUser(id uuid.UUID) (*types.User, error) {
	return nil, nil
}

func (s *userStore) UpdateUser(r *types.UpdateUserRequest) (string, error) {
	return "", nil
}

func (s *userStore) DeleteUser(id uuid.UUID) error {
	return nil
}
