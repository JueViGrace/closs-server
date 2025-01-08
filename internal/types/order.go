package types

type OrderWithLinesResponse struct {
	OrderResponse
	Lines []OrderLineResponse `json:"lines"`
}

type OrderResponse struct {
	KtiNdoc        string  `json:"kti_ndoc"`
	KtiTdoc        string  `json:"kti_tdoc"`
	KtiCodcli      string  `json:"kti_codcli"`
	KtiNombrecli   string  `json:"kti_nombrecli"`
	KtiCodven      string  `json:"kti_codven"`
	KtiDocsol      string  `json:"kti_docsol"`
	KtiCondicion   string  `json:"kti_condicion"`
	KtiTipprec     int     `json:"kti_tipprec"`
	KtiTotneto     float64 `json:"kti_totneto"`
	KtiStatus      string  `json:"kti_status"`
	KtiNroped      string  `json:"kti_nroped"`
	KtiFchdoc      string  `json:"kti_fchdoc"`
	KtiNegesp      bool    `json:"kti_negesp"`
	KePedstatus    string  `json:"ke_pedstatus"`
	Dolarflete     bool    `json:"dolarflete"`
	Complemento    bool    `json:"complemento"`
	NroComplemento string  `json:"nro_complemento"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type OrderLineResponse struct {
	KtiTdoc    string  `json:"kti_tdoc"`
	KtiNdoc    string  `json:"kti_ndoc"`
	KtiTipprec int     `json:"kti_tipprec"`
	KmvCodart  string  `json:"kmv_codart"`
	KmvNombre  string  `json:"kmv_nombre"`
	KmvCant    int     `json:"kmv_cant"`
	KmvArtprec float64 `json:"kmv_artprec"`
	KmvStot    float64 `json:"kmv_stot"`
	KmvDctolin float64 `json:"kmv_dctolin"`
}

type CreateOrderRequest struct {
	KtiNdoc        string  `json:"kti_ndoc"`
	KtiTdoc        string  `json:"kti_tdoc"`
	KtiCodcli      string  `json:"kti_codcli"`
	KtiNombrecli   string  `json:"kti_nombrecli"`
	KtiCodven      string  `json:"kti_codven"`
	KtiDocsol      string  `json:"kti_docsol"`
	KtiCondicion   string  `json:"kti_condicion"`
	KtiTipprec     int     `json:"kti_tipprec"`
	KtiTotneto     float64 `json:"kti_totneto"`
	KtiStatus      string  `json:"kti_status"`
	KtiNroped      string  `json:"kti_nroped"`
	KtiFchdoc      string  `json:"kti_fchdoc"`
	KtiNegesp      bool    `json:"kti_negesp"`
	KePedstatus    string  `json:"ke_pedstatus"`
	Dolarflete     bool    `json:"dolarflete"`
	Complemento    bool    `json:"complemento"`
	NroComplemento string  `json:"nro_complemento"`
}

type CreateOrderLineRequest struct {
	KtiTdoc    string  `json:"kti_tdoc"`
	KtiNdoc    string  `json:"kti_ndoc"`
	KtiTipprec int     `json:"kti_tipprec"`
	KmvCodart  string  `json:"kmv_codart"`
	KmvNombre  string  `json:"kmv_nombre"`
	KmvCant    int     `json:"kmv_cant"`
	KmvArtprec float64 `json:"kmv_artprec"`
	KmvStot    float64 `json:"kmv_stot"`
	KmvDctolin float64 `json:"kmv_dctolin"`
}

type UpdateOrderRequest struct {
	KtiNdoc        string  `json:"kti_ndoc"`
	KtiTdoc        string  `json:"kti_tdoc"`
	KtiCodcli      string  `json:"kti_codcli"`
	KtiNombrecli   string  `json:"kti_nombrecli"`
	KtiCodven      string  `json:"kti_codven"`
	KtiDocsol      string  `json:"kti_docsol"`
	KtiCondicion   string  `json:"kti_condicion"`
	KtiTipprec     int     `json:"kti_tipprec"`
	KtiTotneto     float64 `json:"kti_totneto"`
	KtiStatus      string  `json:"kti_status"`
	KtiNroped      string  `json:"kti_nroped"`
	KtiFchdoc      string  `json:"kti_fchdoc"`
	KtiNegesp      bool    `json:"kti_negesp"`
	KePedstatus    string  `json:"ke_pedstatus"`
	Dolarflete     bool    `json:"dolarflete"`
	Complemento    bool    `json:"complemento"`
	NroComplemento string  `json:"nro_complemento"`
}

type UpdateOrderLineRequest struct {
	KtiTdoc    string  `json:"kti_tdoc"`
	KtiNdoc    string  `json:"kti_ndoc"`
	KtiTipprec int     `json:"kti_tipprec"`
	KmvCodart  string  `json:"kmv_codart"`
	KmvNombre  string  `json:"kmv_nombre"`
	KmvCant    int     `json:"kmv_cant"`
	KmvArtprec float64 `json:"kmv_artprec"`
	KmvStot    float64 `json:"kmv_stot"`
	KmvDctolin float64 `json:"kmv_dctolin"`
}
