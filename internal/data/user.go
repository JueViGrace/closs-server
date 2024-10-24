package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUsers() (users []*types.User, err error)
	GetUser(id *uuid.UUID) (user *types.User, err error)
	UpdateUser(r *types.UpdateUserRequest) (msg string, err error)
	DeleteUser(id *uuid.UUID) (err error)
}

func (s *storage) UserStore() UserStore {
	return NewUserStore(s.ctx, s.queries)
}

type userStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewUserStore(ctx context.Context, db *db.Queries) UserStore {
	return &userStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStore) GetUsers() ([]*types.User, error) {
	users := make([]*types.User, 0)

	dbUsers, err := s.db.AdminGetUsers(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		user, err := types.DbUserToUser(&dbUser)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *userStore) GetUser(id *uuid.UUID) (*types.User, error) {
	user := new(types.User)

	dbUser, err := s.db.GetUserById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	user, err = types.DbUserToUser(&dbUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userStore) UpdateUser(r *types.UpdateUserRequest) (string, error) {
	ur := types.UpdateUserToDb(r)

	err := s.db.UpdateUser(s.ctx, *ur)
	if err != nil {
		return "", err
	}

	return "Updated!", nil
}

func (s *userStore) DeleteUser(id *uuid.UUID) error {
	err := s.db.SoftDeleteUser(s.ctx, id.String())
	if err != nil {
		return err
	}

	return nil
}
