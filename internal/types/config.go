package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Config struct {
	ID           uuid.UUID `json:"id"`
	CnfgIdconfig string    `json:"cnfg_idconfig"`
	CnfgClase    string    `json:"cnfg_clase"`
	CnfgTipo     string    `json:"cnfg_tipo"`
	CnfgValnum   string    `json:"cnfg_valnum"`
	CnfgValsino  bool      `json:"cnfg_valsino"`
	CnfgValtxt   string    `json:"cnfg_valtxt"`
	CnfgLentxt   int16     `json:"cnfg_lentxt"`
	CnfgValfch   time.Time `json:"cnfg_valfch"`
	CnfgActiva   bool      `json:"cnfg_activa"`
	CnfgEtiq     string    `json:"cnfg_etiq"`
	CnfgTtip     string    `json:"cnfg_ttip"`
	UserID       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type CreateConfigRequest struct {
	CnfgIdconfig string    `json:"cnfg_idconfig"`
	CnfgClase    string    `json:"cnfg_clase"`
	CnfgTipo     string    `json:"cnfg_tipo"`
	CnfgValnum   string    `json:"cnfg_valnum"`
	CnfgValsino  bool      `json:"cnfg_valsino"`
	CnfgValtxt   string    `json:"cnfg_valtxt"`
	CnfgLentxt   int16     `json:"cnfg_lentxt"`
	CnfgValfch   time.Time `json:"cnfg_valfch"`
	CnfgActiva   bool      `json:"cnfg_activa"`
	CnfgEtiq     string    `json:"cnfg_etiq"`
	CnfgTtip     string    `json:"cnfg_ttip"`
	UserID       string    `json:"user_id"`
}

type UpdateConfigRequest struct {
	CnfgClase   string    `json:"cnfg_clase"`
	CnfgTipo    string    `json:"cnfg_tipo"`
	CnfgValnum  string    `json:"cnfg_valnum"`
	CnfgValsino bool      `json:"cnfg_valsino"`
	CnfgValtxt  string    `json:"cnfg_valtxt"`
	CnfgLentxt  int16     `json:"cnfg_lentxt"`
	CnfgValfch  time.Time `json:"cnfg_valfch"`
	CnfgActiva  bool      `json:"cnfg_activa"`
	CnfgEtiq    string    `json:"cnfg_etiq"`
	CnfgTtip    string    `json:"cnfg_ttip"`
	ID          uuid.UUID `json:"id"`
}

func DbConfigToConfig(db *db.KeWcnfConf) (*Config, error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	return &Config{
		ID:           id,
		CnfgIdconfig: db.CnfgIdconfig,
		CnfgClase:    db.CnfgClase,
		CnfgTipo:     db.CnfgTipo,
		CnfgValnum:   db.CnfgValnum,
		CnfgValsino:  db.CnfgValsino,
		CnfgValtxt:   db.CnfgValtxt,
		CnfgLentxt:   db.CnfgLentxt,
		CnfgValfch:   db.CnfgValfch,
		CnfgActiva:   db.CnfgActiva,
		CnfgEtiq:     db.CnfgEtiq,
		CnfgTtip:     db.CnfgTtip,
		UserID:       db.UserID,
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		DeletedAt:    db.DeletedAt.Time,
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
