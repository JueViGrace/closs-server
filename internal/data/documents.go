package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
)

type DocumentStore interface {
	GetDocuments()
	GetDocumentsByCode()
	GetDocumentsWithLines()
	GetDocumentsWithLinesByCode()
}

func (s *storage) DocumentStore() DocumentStore {
	return NewDocumentStore(s.ctx, s.queries)
}

type documentStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewDocumentStore(ctx context.Context, db *db.Queries) DocumentStore {
	return &documentStore{
		ctx: ctx,
		db:  db,
	}
}
