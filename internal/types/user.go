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

func DbUserToUser(u *db.Usuario) (*User, error) {
	id, err := uuid.Parse(u.ID)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        id,
		Username:  u.Username,
		Password:  u.Password,
		Role:      Role(u.Role),
		Desactivo: u.Desactivo,
		UltSinc:   u.UltSinc,
		Version:   u.Version,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt.Time,
	}, nil
}

func UpdateUserToDb(r *UpdateUserRequest) *db.UpdateUserParams {
	return &db.UpdateUserParams{
		UltSinc: r.UltSinc,
		Version: r.Version,
		ID:      r.ID.String(),
	}
}
