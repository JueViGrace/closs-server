// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type ClossCompany struct {
	KedCodigo string
	KedNombre string
	KedStatus int64
	KedEnlace string
	KedAgen   string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}

type ClossConfig struct {
	CnfgIdconfig string
	CnfgClase    string
	CnfgTipo     string
	CnfgValnum   float64
	CnfgValsino  int64
	CnfgValtxt   string
	CnfgLentxt   int64
	CnfgValfch   string
	CnfgActiva   int64
	CnfgEtiq     string
	CnfgTtip     string
	Username     string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    sql.NullString
}

type ClossCustomer struct {
	Codigo          string
	Nombre          string
	Email           string
	Direccion       string
	Telefonos       string
	Perscont        string
	Vendedor        string
	Contribespecial int64
	Status          int64
	Sector          string
	Subsector       string
	Precio          int64
	KneActiva       int64
	KneMtomin       float64
	Noemifac        int64
	Noeminota       int64
	Fchultvta       string
	Mtoultvta       float64
	Prcdpagdia      float64
	Promdiasp       float64
	Riesgocrd       float64
	Cantdocs        int64
	Totmtodocs      float64
	Prommtodoc      float64
	Diasultvta      float64
	Promdiasvta     float64
	Limcred         float64
	Fchcrea         string
	Dolarflete      int64
	Nodolarflete    int64
	CreatedAt       string
	UpdatedAt       string
}

type ClossDocument struct {
	Agencia        string
	Tipodoc        string
	Documento      string
	Tipodocv       string
	Codcliente     string
	Nombrecli      string
	Contribesp     int64
	RutaParme      int64
	Tipoprecio     int64
	Emision        string
	Recepcion      string
	Vence          string
	Diascred       float64
	Estatusdoc     int64
	Dtotneto       float64
	Dtotimpuest    float64
	Dtotalfinal    float64
	Dtotpagos      float64
	Dtotdescuen    float64
	Dflete         float64
	Dtotdev        float64
	Dvndmtototal   float64
	Dretencion     float64
	Dretencioniva  float64
	Vendedor       string
	Codcoord       string
	Aceptadev      int64
	KtiNegesp      int64
	Bsiva          float64
	Bsflete        float64
	Bsretencion    float64
	Bsretencioniva float64
	Tasadoc        float64
	Mtodcto        float64
	Fchvencedcto   string
	Tienedcto      int64
	Cbsret         float64
	Cdret          float64
	Cbsretiva      float64
	Cdretiva       float64
	Cbsrparme      float64
	Cdrparme       float64
	Cbsretflete    float64
	Cdretflete     float64
	Bsmtoiva       float64
	Bsmtofte       float64
	RetmunMto      float64
	Dolarflete     int64
	Bsretflete     float64
	Dretflete      float64
	DretmunMto     float64
	Retivaoblig    int64
	Edoentrega     int64
	Dmtoiva        float64
	Prcdctoaplic   float64
	Montodctodol   float64
	Montodctobs    float64
	CreatedAt      string
	UpdatedAt      string
}

type ClossDocumentLine struct {
	Agencia       string
	Tipodoc       string
	Documento     string
	Tipodocv      string
	Grupo         string
	Subgrupo      string
	Origen        int64
	Codigo        string
	Codhijo       string
	Pid           string
	Nombre        string
	Cantidad      int64
	Cntdevuelt    int64
	Vndcntdevuelt int64
	Dvndmtototal  float64
	Dpreciofin    float64
	Dpreciounit   float64
	Dmontoneto    float64
	Dmontototal   float64
	Timpueprc     float64
	Unidevuelt    int64
	Fechadoc      string
	Vendedor      string
	Codcoord      string
}

type ClossGroup struct {
	Codigo    string
	Nombre    string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}

type ClossManager struct {
	KngCodgcia  string
	KngCodcoord string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   sql.NullString
}

type ClossOrder struct {
	KtiNdoc        string
	KtiTdoc        string
	KtiCodcli      string
	KtiNombrecli   string
	KtiCodven      string
	KtiDocsol      string
	KtiCondicion   string
	KtiTipprec     int64
	KtiTotneto     string
	KtiStatus      int64
	KtiNroped      string
	KtiFchdoc      string
	KtiNegesp      int64
	KePedstatus    int64
	Dolarflete     int64
	Complemento    int64
	NroComplemento string
	CreatedAt      string
	UpdatedAt      string
}

type ClossOrderLine struct {
	KtiTdoc    string
	KtiNdoc    string
	KtiTipprec string
	KmvCodart  string
	KmvNombre  string
	KmvCant    int64
	KmvArtprec float64
	KmvStot    float64
	KmvDctolin float64
	CreatedAt  string
	UpdatedAt  string
}

type ClossProduct struct {
	Codigo       string
	Grupo        string
	Subgrupo     string
	Nombre       string
	Referencia   string
	Marca        string
	Unidad       string
	Discont      int64
	Existencia   int64
	VtaMax       int64
	VtaMin       int64
	VtaMinenx    int64
	Comprometido int64
	Precio1      float64
	Precio2      float64
	Precio3      float64
	Precio4      float64
	Precio5      float64
	Precio6      float64
	Precio7      float64
	Preventa     int64
	Dctotope     float64
	VtaSolofac   int64
	VtaSolone    int64
	Codbarras    int64
	Detalles     string
	Cantbulto    float64
	CostoProm    float64
	Util1        string
	Util2        string
	Util3        string
	Fchultcomp   string
	Qtyultcomp   string
	Images       string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    sql.NullString
}

type ClossSalesman struct {
	Codigo    string
	Nombre    string
	Email     string
	Telefono  string
	Telefonos string
	Status    int64
	Supervpor string
	Sector    string
	Subcodigo string
	Nivgcial  int64
	CreatedAt string
	UpdatedAt string
}

type ClossSalesmanStatistic struct {
	Codcoord      string
	Nomcoord      string
	Vendedor      string
	Nombrevend    string
	Cntpedidos    int64
	Mtopedidos    float64
	Cntfacturas   int64
	Mtofacturas   float64
	Metavend      float64
	Prcmeta       float64
	Cntclientes   int64
	Clivisit      int64
	Prcvisitas    float64
	LomMontovtas  float64
	LomPrcvtas    float64
	LomPrcvisit   float64
	RlomMontovtas float64
	RlomPrcvtas   float64
	RlomPrcvisit  float64
	FechaEstad    string
	PpgdolTotneto float64
	DevdolTotneto float64
	DefdolTotneto float64
	Totdolcob     float64
	Cntrecl       int64
	Mtorecl       float64
	CreatedAt     string
	UpdatedAt     string
}

type ClossSector struct {
	Codigo    string
	Zona      string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}

type ClossSession struct {
	Token  string
	UserID string
}

type ClossSubgroup struct {
	Codigo    string
	Subcodigo string
	Nombre    string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}

type ClossSubsector struct {
	Codigo    string
	Subcodigo string
	Subsector string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}

type ClossUser struct {
	ID        string
	Username  string
	Password  string
	Codigo    string
	Role      string
	UltSinc   string
	Version   string
	CreatedAt string
	UpdatedAt string
	DeletedAt sql.NullString
}
