package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	Desactivo bool      `json:"desactivo"`
	UltSinc   time.Time `json:"ult_sinc"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UpdateUserRequest struct {
	UltSinc time.Time `json:"ult_sinc"`
	Version string    `json:"version"`
	ID      uuid.UUID `json:"id"`
}

type Role string

const (
	RoleCliente  Role = "cliente"
	RoleVendedor Role = "vendedor"
	RoleGerente  Role = "gerente"
	RoleAdmin    Role = "administrador"
)

func DbUserToUser(db *db.Usuario) (*User, error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        id,
		Username:  db.Username,
		Password:  db.Password,
		Role:      Role(db.Role),
		Desactivo: db.Desactivo,
		UltSinc:   db.UltSinc,
		Version:   db.Version,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		DeletedAt: db.DeletedAt.Time,
	}, nil
}

func UpdateUserToDb(r *UpdateUserRequest) *db.UpdateUserParams {
	return &db.UpdateUserParams{
		UltSinc: r.UltSinc,
		Version: r.Version,
		ID:      r.ID.String(),
	}
}
