package types

import "github.com/JueViGrace/clo-backend/internal/db"

type Statistic struct{}

type CreateStatisticRequest struct{}

type UpdateStatisticRequest struct{}

func DbStatisticToStatistc(db *db.KeEstadc01) (*Statistic, error) {
	return &Statistic{}, nil
}

func CreateStatisticToDb(r *CreateStatisticRequest) (*db.CreateStatisticParams, error) {
	return &db.CreateStatisticParams{}, nil
}

func UpdateStatisticToDb(r *UpdateStatisticRequest) *db.UpdateStatisticParams {
	return &db.UpdateStatisticParams{}
}
