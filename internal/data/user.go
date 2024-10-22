package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers()
	GetUser(id uuid.UUID)
	CreateUser()
	UpdateUser(id uuid.UUID)
	DeleteUser(id uuid.UUID)
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

func (s *userStore) GetUsers() {
}

func (s *userStore) GetUser(id uuid.UUID) {
}

func (s *userStore) CreateUser() {
}

func (s *userStore) UpdateUser(id uuid.UUID) {
}

func (s *userStore) DeleteUser(id uuid.UUID) {
}
