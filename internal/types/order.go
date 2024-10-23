package types

import (
	"time"

	"github.com/google/uuid"
)

type OrderWithLines struct {
	Order
	Lines []OrderLine `json:"lines"`
}

type Order struct {
	ID             uuid.UUID `json:"id"`
	KtiNdoc        string    `json:"kti_ndoc"`
	KtiTdoc        string    `json:"kti_tdoc"`
	KtiCodcli      string    `json:"kti_codcli"`
	KtiNombrecli   string    `json:"kti_nombrecli"`
	KtiCodven      string    `json:"kti_codven"`
	KtiDocsol      string    `json:"kti_docsol"`
	KtiCondicion   string    `json:"kti_condicion"`
	KtiTipprec     string    `json:"kti_tipprec"`
	KtiTotneto     string    `json:"kti_totneto"`
	KtiStatus      string    `json:"kti_status"`
	KtiNroped      string    `json:"kti_nroped"`
	KtiFchdoc      time.Time `json:"kti_fchdoc"`
	KtiNegesp      bool      `json:"kti_negesp"`
	KePedstatus    string    `json:"ke_pedstatus"`
	Dolarflete     bool      `json:"dolarflete"`
	Complemento    bool      `json:"complemento"`
	NroComplemento string    `json:"nro_complemento"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type OrderLine struct {
	PedidoID   uuid.UUID `json:"pedido_id"`
	ArticuloID uuid.UUID `json:"articulo_id"`
	KtiTdoc    string    `json:"kti_tdoc"`
	KtiNdoc    string    `json:"kti_ndoc"`
	KtiTipprec string    `json:"kti_tipprec"`
	KmvCodart  string    `json:"kmv_codart"`
	KmvNombre  string    `json:"kmv_nombre"`
	KmvCant    int32     `json:"kmv_cant"`
	KmvArtprec string    `json:"kmv_artprec"`
	KmvStot    string    `json:"kmv_stot"`
	KmvDctolin string    `json:"kmv_dctolin"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type CreateOrderRequest struct {
	KtiNdoc        string    `json:"kti_ndoc"`
	KtiTdoc        string    `json:"kti_tdoc"`
	KtiCodcli      string    `json:"kti_codcli"`
	KtiNombrecli   string    `json:"kti_nombrecli"`
	KtiCodven      string    `json:"kti_codven"`
	KtiDocsol      string    `json:"kti_docsol"`
	KtiCondicion   string    `json:"kti_condicion"`
	KtiTipprec     string    `json:"kti_tipprec"`
	KtiTotneto     string    `json:"kti_totneto"`
	KtiStatus      string    `json:"kti_status"`
	KtiNroped      string    `json:"kti_nroped"`
	KtiFchdoc      time.Time `json:"kti_fchdoc"`
	KtiNegesp      bool      `json:"kti_negesp"`
	KePedstatus    string    `json:"ke_pedstatus"`
	Dolarflete     bool      `json:"dolarflete"`
	Complemento    bool      `json:"complemento"`
	NroComplemento string    `json:"nro_complemento"`
}

type CreateOrderLineRequest struct {
	PedidoID   uuid.UUID `json:"pedido_id"`
	ArticuloID uuid.UUID `json:"articulo_id"`
	KtiTdoc    string    `json:"kti_tdoc"`
	KtiNdoc    string    `json:"kti_ndoc"`
	KtiTipprec string    `json:"kti_tipprec"`
	KmvCodart  string    `json:"kmv_codart"`
	KmvNombre  string    `json:"kmv_nombre"`
	KmvCant    int32     `json:"kmv_cant"`
	KmvArtprec string    `json:"kmv_artprec"`
	KmvStot    string    `json:"kmv_stot"`
	KmvDctolin string    `json:"kmv_dctolin"`
}

type UpdateOrderRequest struct {
	KtiTdoc        string    `json:"kti_tdoc"`
	KtiCodcli      string    `json:"kti_codcli"`
	KtiNombrecli   string    `json:"kti_nombrecli"`
	KtiCodven      string    `json:"kti_codven"`
	KtiDocsol      string    `json:"kti_docsol"`
	KtiCondicion   string    `json:"kti_condicion"`
	KtiTipprec     string    `json:"kti_tipprec"`
	KtiTotneto     string    `json:"kti_totneto"`
	KtiStatus      string    `json:"kti_status"`
	KtiNroped      string    `json:"kti_nroped"`
	KtiFchdoc      time.Time `json:"kti_fchdoc"`
	KtiNegesp      bool      `json:"kti_negesp"`
	KePedstatus    string    `json:"ke_pedstatus"`
	Dolarflete     bool      `json:"dolarflete"`
	Complemento    bool      `json:"complemento"`
	NroComplemento string    `json:"nro_complemento"`
}

type UpdateOrderLineRequest struct {
	KtiTdoc    string `json:"kti_tdoc"`
	KtiTipprec string `json:"kti_tipprec"`
	KmvNombre  string `json:"kmv_nombre"`
	KmvCant    int32  `json:"kmv_cant"`
	KmvArtprec string `json:"kmv_artprec"`
	KmvStot    string `json:"kmv_stot"`
	KmvDctolin string `json:"kmv_dctolin"`
	PedidoID   string `json:"pedido_id"`
	ArticuloID string `json:"articulo_id"`
}
