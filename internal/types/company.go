package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID
	KedCodigo string
	KedNombre string
	KedStatus bool
	KedEnlace string
	KedAgen   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreateCompanyRequest struct {
	KedCodigo string
	KedNombre string
	KedStatus bool
	KedEnlace string
	KedAgen   string
}

type UpdateCompanyRequest struct {
	KedNombre string
	KedStatus bool
	KedEnlace string
	KedAgen   string
	ID        uuid.UUID
}

func DbComanyToCompany(com *db.KeDataconex) (*Company, error) {
	id, err := uuid.Parse(com.ID)
	if err != nil {
		return nil, err
	}

	return &Company{
		ID:        id,
		KedCodigo: com.KedCodigo,
		KedNombre: com.KedNombre,
		KedStatus: com.KedStatus,
		KedEnlace: com.KedEnlace,
		KedAgen:   com.KedAgen,
		CreatedAt: com.CreatedAt,
		UpdatedAt: com.UpdatedAt,
		DeletedAt: com.DeletedAt.Time,
	}, nil
}

func CreateCompanyToDb(r *CreateCompanyRequest) (*db.CreateCompanyParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &db.CreateCompanyParams{
		ID:        id.String(),
		KedCodigo: r.KedCodigo,
		KedNombre: r.KedNombre,
		KedStatus: r.KedStatus,
		KedEnlace: r.KedEnlace,
		KedAgen:   r.KedAgen,
	}, nil
}

func UpdateCompanyToDb(r *UpdateCompanyRequest) *db.UpdateCompanyParams {
	return &db.UpdateCompanyParams{
		KedNombre: r.KedNombre,
		KedStatus: r.KedStatus,
		KedEnlace: r.KedEnlace,
		KedAgen:   r.KedAgen,
		ID:        r.ID.String(),
	}
}
