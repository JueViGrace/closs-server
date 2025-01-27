package handlers

import (
	"github.com/JueViGrace/clo-backend/internal/data"
	"github.com/JueViGrace/clo-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type DocumentHandler interface {
	GetDocuments(c *fiber.Ctx, d *types.AuthData) error
	GetDocumentByCode(c *fiber.Ctx, d *types.AuthData) error
	CreateDocument(c *fiber.Ctx) error
	UpdateDocument(c *fiber.Ctx) error
}

type documentHandler struct {
	db data.DocumentStore
}

func NewDocumentHandler(db data.DocumentStore) DocumentHandler {
	return &documentHandler{
		db: db,
	}
}

// todo: use auth data for search
func (h *documentHandler) GetDocuments(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)
	mq := new(types.ManagerDocumentQueries)
	if err := c.QueryParser(mq); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	sq := new(types.SalesmanDocumentQueries)
	if err := c.QueryParser(sq); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	cq := new(types.CustomerDocumentQueries)
	if err := c.QueryParser(cq); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	switch {
	case mq != nil:
		res = types.RespondOk(nil, "manager handler")
	case sq != nil:

		res = types.RespondOk(nil, "salesman handler")
	case cq != nil:

		res = types.RespondOk(nil, "customer handler")
	case c.QueryBool("withLines", false):
		docs, err := h.db.GetDocumentsWithLines()
		if err != nil {
			res = types.RespondNotFound(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
		res = types.RespondOk(docs, "Success")

	default:
		docs, err := h.db.GetDocuments()
		if err != nil {
			res = types.RespondNotFound(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}
		res = types.RespondOk(docs, "Success")
	}

	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentByCode(c *fiber.Ctx, d *types.AuthData) error {
	res := new(types.APIResponse)

	switch {
	case c.QueryBool("withLines", false):
		docs, err := h.db.GetDocumentByCodeWithLines(c.Params("code"))
		if err != nil {
			res = types.RespondNotFound(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}

		res = types.RespondOk(docs, "Success")

	default:
		docs, err := h.db.GetDocumentByCode(c.Params("code"))
		if err != nil {
			res = types.RespondNotFound(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}

		res = types.RespondOk(docs, "Success")
	}
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsByManager(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	docs, err := h.db.GetDocumentsByManager(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsBySalesman(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	docs, err := h.db.GetDocumentsBySalesman(c.Params("code"))
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) GetDocumentsByCustomer(c *fiber.Ctx) error {
	res := new(types.APIResponse)

	docs, err := h.db.GetDocumentsByCustomer(c.Params("customer"))
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondOk(docs, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) CreateDocument(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.CreateDocumentRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.CreateDocument(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}

func (h *documentHandler) UpdateDocument(c *fiber.Ctx) error {
	res := new(types.APIResponse)
	r := new(types.UpdateDocumentRequest)
	if err := c.BodyParser(r); err != nil {
		res = types.RespondBadRequest(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	m, err := h.db.UpdateDocument(r)
	if err != nil {
		res = types.RespondNotFound(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res = types.RespondCreated(m, "Success")
	return c.Status(res.Status).JSON(res)
}
