package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
	"github.com/google/uuid"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Version  string `json:"version"`
}

func CreateUserToDb(r *SignUpRequest) (*db.CreateUserParams, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	// TODO: hash password
	pass, err := util.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	return &db.CreateUserParams{
		ID:        id.String(),
		Username:  r.Username,
		Password:  pass,
		Role:      db.UsuarioRole(Role("cliente")),
		Desactivo: false,
		UltSinc:   time.Now().UTC(),
		Version:   r.Version,
	}, nil
}
