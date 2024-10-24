package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type DocumentStore interface {
	GetDocuments() (documents []*types.Document, err error)
	GetDocumentsBySalesman(code string) (documents []*types.Document, err error)
	GetDocumentsWithLines() (documents []*types.DocumentWithLines, err error)
	GetDocumentsWithLinesBySalesman(code string) (documents []*types.DocumentWithLines, err error)
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

	dbDocs, err := s.db.AdminGetDocuments(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		doc, err := types.DbKeDoccToDocument(&dbDoc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsBySalesman(code string) ([]*types.Document, error) {
	docs := make([]*types.Document, 0)

	dbDocs, err := s.db.GetDocumentsBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		doc, err := types.DbKeDoccToDocument(&dbDoc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsWithLines() ([]*types.DocumentWithLines, error) {
	docs := make([]*types.DocumentWithLines, 0)
	docMap := make(map[types.Document][]types.DocumentLine)
	doc := new(types.Document)

	dbDocs, err := s.db.AdminGetDocumentsWithLines(s.ctx)
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

func (s *documentStore) GetDocumentsWithLinesBySalesman(code string) ([]*types.DocumentWithLines, error) {
	docs := make([]*types.DocumentWithLines, 0)
	docMap := make(map[types.Document][]types.DocumentLine)
	doc := new(types.Document)

	dbDocs, err := s.db.GetDocumentsWithLinesBySalesman(s.ctx, code)
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
