package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/google/uuid"
)

type SessionStore interface {
	GetSessionById(id uuid.UUID) (*types.Session, error)
	GetSessionByUsername(username string) (*types.Session, error)
	DeleteSessionById(id uuid.UUID) error
}

func (s *storage) SessionStore() SessionStore {
	return NewSessionStore(s.ctx, s.queries)
}

type sessionStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewSessionStore(ctx context.Context, db *db.Queries) SessionStore {
	return &sessionStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *sessionStore) GetSessionById(id uuid.UUID) (*types.Session, error) {
	session, err := s.db.GetSessionById(s.ctx, id.String())
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(session.UserID)
	if err != nil {
		return nil, err
	}

	return &types.Session{
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
		Username:     session.Username,
		UserId:       userId,
	}, nil
}

func (s *sessionStore) GetSessionByUsername(username string) (*types.Session, error) {
	session, err := s.db.GetSessionByUsername(s.ctx, username)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(session.UserID)
	if err != nil {
		return nil, err
	}

	return &types.Session{
		RefreshToken: session.RefreshToken,
		AccessToken:  session.AccessToken,
		Username:     session.Username,
		UserId:       userId,
	}, nil
}

func (s *sessionStore) DeleteSessionById(id uuid.UUID) error {
	err := s.db.DeleteSessionById(s.ctx, id.String())
	if err != nil {
		return err
	}
	return nil
}
