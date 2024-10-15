package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type DocumentStore interface {
	GetDocuments() ([]*types.Document, error)
	GetDocumentsByCode(code string) ([]*types.Document, error)
	GetDocumentsWithLines() ([]*types.DocumentWithLines, error)
	GetDocumentsWithLinesByCode(code string) ([]*types.DocumentWithLines, error)
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

func (s *documentStore) GetDocuments() ([]*types.Document, error) {
	docs := make([]*types.Document, 0)

	dbDocs, err := s.db.FindAllDocuments(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		docs = append(docs, types.DbKeDoccToDocument(&dbDoc))
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsByCode(code string) ([]*types.Document, error) {
	docs := make([]*types.Document, 0)

	dbDocs, err := s.db.FindAllDocumentsByCode(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		docs = append(docs, types.DbKeDoccToDocument(&dbDoc))
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsWithLines() ([]*types.DocumentWithLines, error) {
	docs := make([]*types.DocumentWithLines, 0)
	docMap := make(map[types.Document][]types.DocumentLine)
	doc := new(types.Document)

	dbDocs, err := s.db.FindAllDocumentsWithLines(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		if doc == nil {
			doc = types.DbDocToDocument(&dbDoc)
		}

		if doc.Documento != dbDoc.Documento {
			doc = types.DbDocToDocument(&dbDoc)
		}

		docMap[*doc] = append(docMap[*doc], *types.DbDocToDocLine(&dbDoc))
	}

	for key, value := range docMap {
		docs = append(docs, types.DocMapToDocWithLines(&key, &value))
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsWithLinesByCode(code string) ([]*types.DocumentWithLines, error) {
	docs := make([]*types.DocumentWithLines, 0)
	docMap := make(map[types.Document][]types.DocumentLine)
	doc := new(types.Document)

	dbDocs, err := s.db.FindAllDocumentsWithLinesByCode(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		if doc == nil {
			doc = types.DbDocByCodeToDocument(&dbDoc)
		}

		if doc.Documento != dbDoc.Documento {
			doc = types.DbDocByCodeToDocument(&dbDoc)
		}

		docMap[*doc] = append(docMap[*doc], *types.DbDocByCodeToDocLine(&dbDoc))
	}

	for key, value := range docMap {
		docs = append(docs, types.DocMapToDocWithLines(&key, &value))
	}

	return docs, nil
}
