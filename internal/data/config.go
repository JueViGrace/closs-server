package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type ConfigStore interface {
	GetCofigs() ([]*types.Config, error)
	GetConfig(id uuid.UUID) (*types.Config, error)
	CreateConfig(r *types.CreateConfigRequest) (string, error)
	UpdateConfig(r *types.UpdateConfigRequest) (string, error)
	DeleteConfig(id uuid.UUID) error
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

func (s *configStore) GetCofigs() ([]*types.Config, error) {
	return nil, nil
}

func (s *configStore) GetConfig(id uuid.UUID) (*types.Config, error) {
	return nil, nil
}

func (s *configStore) CreateConfig(r *types.CreateConfigRequest) (string, error) {
	return "", nil
}

func (s *configStore) UpdateConfig(r *types.UpdateConfigRequest) (string, error) {
	return "", nil
}

func (s *configStore) DeleteConfig(id uuid.UUID) error {
	return nil
}
