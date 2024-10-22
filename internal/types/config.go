package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Config struct {
	ID           uuid.UUID
	CnfgIdconfig string
	CnfgClase    string
	CnfgTipo     string
	CnfgValnum   string
	CnfgValsino  bool
	CnfgValtxt   string
	CnfgLentxt   int16
	CnfgValfch   time.Time
	CnfgActiva   bool
	CnfgEtiq     string
	CnfgTtip     string
	UserID       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type CreateConfigRequest struct {
	CnfgIdconfig string
	CnfgClase    string
	CnfgTipo     string
	CnfgValnum   string
	CnfgValsino  bool
	CnfgValtxt   string
	CnfgLentxt   int16
	CnfgValfch   time.Time
	CnfgActiva   bool
	CnfgEtiq     string
	CnfgTtip     string
	UserID       string
}

type UpdateConfigRequest struct {
	CnfgClase   string
	CnfgTipo    string
	CnfgValnum  string
	CnfgValsino bool
	CnfgValtxt  string
	CnfgLentxt  int16
	CnfgValfch  time.Time
	CnfgActiva  bool
	CnfgEtiq    string
	CnfgTtip    string
	ID          uuid.UUID
}

func DbConfigToConfig(cfg *db.KeWcnfConf) (*Config, error) {
	id, err := uuid.Parse(cfg.ID)
	if err != nil {
		return nil, err
	}

	return &Config{
		ID:           id,
		CnfgIdconfig: cfg.CnfgIdconfig,
		CnfgClase:    cfg.CnfgClase,
		CnfgTipo:     cfg.CnfgTipo,
		CnfgValnum:   cfg.CnfgValnum,
		CnfgValsino:  cfg.CnfgValsino,
		CnfgValtxt:   cfg.CnfgValtxt,
		CnfgLentxt:   cfg.CnfgLentxt,
		CnfgValfch:   cfg.CnfgValfch,
		CnfgActiva:   cfg.CnfgActiva,
		CnfgEtiq:     cfg.CnfgEtiq,
		CnfgTtip:     cfg.CnfgTtip,
		UserID:       cfg.UserID,
		CreatedAt:    cfg.CreatedAt,
		UpdatedAt:    cfg.UpdatedAt,
		DeletedAt:    cfg.DeletedAt.Time,
	}, nil
}

func CreateConfigToDb(r *CreateConfigRequest) (*db.CreateConfigParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &db.CreateConfigParams{
		ID:           id.String(),
		CnfgIdconfig: r.CnfgIdconfig,
		CnfgClase:    r.CnfgClase,
		CnfgTipo:     r.CnfgTipo,
		CnfgValnum:   r.CnfgValnum,
		CnfgValsino:  r.CnfgValsino,
		CnfgValtxt:   r.CnfgValtxt,
		CnfgLentxt:   r.CnfgLentxt,
		CnfgValfch:   r.CnfgValfch,
		CnfgActiva:   r.CnfgActiva,
		CnfgEtiq:     r.CnfgEtiq,
		CnfgTtip:     r.CnfgTtip,
		UserID:       r.UserID,
	}, nil
}

func UpdateConfigToDb(r *UpdateConfigRequest) *db.UpdateConfigParams {
	return &db.UpdateConfigParams{
		CnfgClase:   r.CnfgClase,
		CnfgTipo:    r.CnfgTipo,
		CnfgValnum:  r.CnfgValnum,
		CnfgValsino: r.CnfgValsino,
		CnfgValtxt:  r.CnfgValtxt,
		CnfgLentxt:  r.CnfgLentxt,
		CnfgValfch:  r.CnfgValfch,
		CnfgActiva:  r.CnfgActiva,
		CnfgEtiq:    r.CnfgEtiq,
		CnfgTtip:    r.CnfgTtip,
		ID:          r.ID.String(),
	}
}
