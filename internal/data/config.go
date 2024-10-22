package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type ConfigStore interface {
	GetCofigs()
	GetConfig(id uuid.UUID)
	CreateConfig()
	UpdateConfig(id uuid.UUID)
	DeleteConfig(id uuid.UUID)
}

func (s *storage) ConfigStore() ConfigStore {
	return NewConfigStore(s.ctx, s.queries)
}

type configStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewConfigStore(ctx context.Context, db *db.Queries) ConfigStore {
	return &configStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *configStore) GetCofigs() {
}

func (s *configStore) GetConfig(id uuid.UUID) {
}

func (s *configStore) CreateConfig() {
}

func (s *configStore) UpdateConfig(id uuid.UUID) {
}

func (s *configStore) DeleteConfig(id uuid.UUID) {
}
