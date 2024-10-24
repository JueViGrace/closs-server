package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type StatisticStore interface {
	GetStatistics() ([]*types.Statistic, error)
	GetStatistic(id *uuid.UUID) (*types.Statistic, error)
	CreateStatistic(r *types.CreateStatisticRequest) (string, error)
	UpdateStatistic(r *types.UpdateStatisticRequest) (string, error)
	DeleteStatistic(id *uuid.UUID) error
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
	statistics := make([]*types.Statistic, 0)

	dbStatistics, err := s.db.AdminGetStatistics(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbStatistic := range dbStatistics {
		statistic, err := types.DbStatisticToStatistc(&dbStatistic)
		if err != nil {
			return nil, err
		}
		statistics = append(statistics, statistic)
	}

	return statistics, nil
}

func (s *statisticStore) GetStatistic(id *uuid.UUID) (*types.Statistic, error) {
	statistic := new(types.Statistic)

	dbStatistic, err := s.db.GetStatisticsById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	statistic, err = types.DbStatisticToStatistc(&dbStatistic)
	if err != nil {
		return nil, err
	}

	return statistic, nil
}

func (s *statisticStore) CreateStatistic(r *types.CreateStatisticRequest) (string, error) {
	cr, err := types.CreateStatisticToDb(r)
	if err != nil {
		return "", err
	}

	err = s.db.CreateStatistic(s.ctx, *cr)
	if err != nil {
		return "", err
	}

	return "Created!", nil
}

func (s *statisticStore) UpdateStatistic(r *types.UpdateStatisticRequest) (string, error) {
	ur := types.UpdateStatisticToDb(r)

	err := s.db.UpdateStatistic(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updated!", nil
}

func (s *statisticStore) DeleteStatistic(id *uuid.UUID) error {
	err := s.db.SoftDeleteStatistic(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}
