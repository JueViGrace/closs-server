package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type ConfigStore interface {
	AdminGetCofigs() (configs []*types.ConfigResponse, err error)
	AdminGetConfigById(id *uuid.UUID) (config *types.ConfigResponse, err error)
	CreateConfig(r *types.CreateConfigRequest) (msg string, err error)
	UpdateConfig(r *types.UpdateConfigRequest) (msg string, err error)
	DeleteConfig(id *uuid.UUID) (err error)
	GetConfigByUser(userId *uuid.UUID) (configs []*types.ConfigResponse, err error)
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

func (s *configStore) AdminGetCofigs() ([]*types.ConfigResponse, error) {
	configs := make([]*types.ConfigResponse, 0)

	dbConfigs, err := s.db.AdminGetConfigs(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbConfig := range dbConfigs {
		config, err := types.DbConfigToConfig(&dbConfig)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func (s *configStore) AdminGetConfigById(id *uuid.UUID) (*types.ConfigResponse, error) {
	config := new(types.ConfigResponse)

	dbConfig, err := s.db.AdminGetConfigById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	config, err = types.DbConfigToConfig(&dbConfig)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (s *configStore) CreateConfig(r *types.CreateConfigRequest) (string, error) {
	cr, err := types.CreateConfigToDb(r)
	if err != nil {
		return "", err
	}

	err = s.db.CreateConfig(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *configStore) UpdateConfig(r *types.UpdateConfigRequest) (string, error) {
	ur := types.UpdateConfigToDb(r)

	err := s.db.UpdateConfig(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updated!", nil
}

func (s *configStore) DeleteConfig(id *uuid.UUID) error {
	err := s.db.SoftDeleteConfig(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}

func (s *configStore) GetConfigByUser(userId *uuid.UUID) ([]*types.ConfigResponse, error) {
	configs := make([]*types.ConfigResponse, 0)

	dbConfigs, err := s.db.GetConfigByUser(s.ctx, userId.String())
	if err != nil {
		return nil, err
	}

	for _, dbConfig := range dbConfigs {
		config, err := types.DbConfigToConfig(&dbConfig)
		if err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}

	return configs, nil
}
