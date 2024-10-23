package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type StatisticStore interface {
	GetStatistics() ([]*types.Statistic, error)
	GetStatistic(id uuid.UUID) (*types.Statistic, error)
	CreateStatistic(r *types.CreateStatisticRequest) (string, error)
	UpdateStatistic(r *types.UpdateStatisticRequest) (string, error)
	DeleteStatistic(id uuid.UUID) error
}

func (s *storage) StatisticStore() StatisticStore {
	return NewStatisticStore(s.ctx, s.queries)
}

type statisticStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewStatisticStore(ctx context.Context, db *db.Queries) StatisticStore {
	return &statisticStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *statisticStore) GetStatistics() ([]*types.Statistic, error) {
	return nil, nil
}

func (s *statisticStore) GetStatistic(id uuid.UUID) (*types.Statistic, error) {
	return nil, nil
}

func (s *statisticStore) CreateStatistic(r *types.CreateStatisticRequest) (string, error) {
	return "", nil
}

func (s *statisticStore) UpdateStatistic(r *types.UpdateStatisticRequest) (string, error) {
	return "", nil
}

func (s *statisticStore) DeleteStatistic(id uuid.UUID) error {
	return nil
}
