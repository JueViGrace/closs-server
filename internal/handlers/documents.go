package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type DocumentHandler interface {
	GetDocuments(c *fiber.Ctx) error
	GetDocumentsByCode(c *fiber.Ctx) error
	GetDocumentsWithLines(c *fiber.Ctx) error
	GetDocumentsWithLinesByCode(c *fiber.Ctx) error
}

type documentHandler struct {
	db data.DocumentStore
}

func NewDocumentHandler(db data.DocumentStore) DocumentHandler {
	return &documentHandler{
		db: db,
	}
}

func (h *documentHandler) GetDocuments(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	docs, err := h.db.GetDocuments()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed to find documents")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id := c.Params("id")
	docs, err := h.db.GetDocumentsByCode(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed to find documents")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsWithLines(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	docs, err := h.db.GetDocumentsWithLines()
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed to find documents")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsWithLinesByCode(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	id := c.Params("id")
	docs, err := h.db.GetDocumentsWithLinesByCode(id)
	if err != nil {
		res = types.RespondNotFound(err.Error(), "Failed to find documents")
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}
