package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/util"
)

type DocumentWithLinesResponse struct {
	DocumentResponse
	Lines []DocumentLineResponse `json:"lines"`
}

type DocumentResponse struct {
	Agencia        string  `json:"agencia"`
	Tipodoc        string  `json:"tipodoc"`
	Documento      string  `json:"documento"`
	Tipodocv       string  `json:"tipodocv"`
	Codcliente     string  `json:"codcliente"`
	Nombrecli      string  `json:"nombrecli"`
	Contribesp     bool    `json:"contribesp"`
	RutaParme      bool    `json:"ruta_parme"`
	Tipoprecio     int     `json:"tipoprecio"`
	Emision        string  `json:"emision"`
	Recepcion      string  `json:"recepcion"`
	Vence          string  `json:"vence"`
	Diascred       int     `json:"diascred"`
	Estatusdoc     int     `json:"estatusdoc"`
	Dtotneto       float64 `json:"dtotneto"`
	Dtotimpuest    float64 `json:"dtotimpuest"`
	Dtotalfinal    float64 `json:"dtotalfinal"`
	Dtotpagos      float64 `json:"dtotpagos"`
	Dtotdescuen    float64 `json:"dtotdescuen"`
	Dflete         float64 `json:"dFlete"`
	Dtotdev        float64 `json:"dtotdev"`
	Dvndmtototal   float64 `json:"dvndmtototal"`
	Dretencion     float64 `json:"dretencion"`
	Dretencioniva  float64 `json:"dretencioniva"`
	Vendedor       string  `json:"vendedor"`
	Codcoord       string  `json:"codcoord"`
	Aceptadev      bool    `json:"aceptadev"`
	KtiNegesp      bool    `json:"kti_negesp"`
	Bsiva          float64 `json:"bsiva"`
	Bsflete        float64 `json:"bsflete"`
	Bsretencion    float64 `json:"bsretencion"`
	Bsretencioniva float64 `json:"bsretencioniva"`
	Tasadoc        float64 `json:"tasadoc"`
	Mtodcto        float64 `json:"mtodcto"`
	Fchvencedcto   string  `json:"fchvencedcto"`
	Tienedcto      bool    `json:"tienedcto"`
	Cbsret         float64 `json:"cbsret"`
	Cdret          float64 `json:"cdret"`
	Cbsretiva      float64 `json:"cbsretiva"`
	Cdretiva       float64 `json:"cdretiva"`
	Cbsrparme      float64 `json:"cbsrparme"`
	Cdrparme       float64 `json:"cdrparme"`
	Cbsretflete    float64 `json:"cbsretflete"`
	Cdretflete     float64 `json:"cdretflete"`
	Bsmtoiva       float64 `json:"bsmtoiva"`
	Bsmtofte       float64 `json:"bsmtofte"`
	RetmunMto      float64 `json:"retmun_mto"`
	Dolarflete     bool    `json:"dolarflete"`
	Bsretflete     float64 `json:"bsretflete"`
	Dretflete      float64 `json:"dretflete"`
	DretmunMto     float64 `json:"dretmun_mto"`
	Retivaoblig    bool    `json:"retivaoblig"`
	Edoentrega     bool    `json:"edoentrega"`
	Dmtoiva        float64 `json:"dmtoiva"`
	Prcdctoaplic   float64 `json:"prcdctoaplic"`
	Montodctodol   float64 `json:"montodctodol"`
	Montodctobs    float64 `json:"montodctobs"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type DocumentLineResponse struct {
	Agencia       string  `json:"agencia"`
	Tipodoc       string  `json:"tipodoc"`
	Documento     string  `json:"documento"`
	Tipodocv      string  `json:"tipodocv"`
	Grupo         string  `json:"grupo"`
	Subgrupo      string  `json:"subgrupo"`
	Origen        int     `json:"origen"`
	Codigo        string  `json:"codigo"`
	Codhijo       string  `json:"codhijo"`
	Pid           string  `json:"pid"`
	Nombre        string  `json:"nombre"`
	Cantidad      int     `json:"cantidad"`
	Cntdevuelt    int     `json:"cntdevuelt"`
	Vndcntdevuelt int     `json:"vndcntdevuelt"`
	Dvndmtototal  float64 `json:"dvndmtototal"`
	Dpreciofin    float64 `json:"dpreciofin"`
	Dpreciounit   float64 `json:"dpreciounit"`
	Dmontoneto    float64 `json:"dmontoneto"`
	Dmontototal   float64 `json:"dmontototal"`
	Timpueprc     float64 `json:"timpueprc"`
	Unidevuelt    int     `json:"unidevuelt"`
	Fechadoc      string  `json:"fechadoc"`
	Vendedor      string  `json:"vendedor"`
	Codcoord      string  `json:"codcoord"`
}

type CreateDocumentRequest struct {
	Agencia        string                      `json:"agencia"`
	Tipodoc        string                      `json:"tipodoc"`
	Documento      string                      `json:"documento"`
	Tipodocv       string                      `json:"tipodocv"`
	Codcliente     string                      `json:"codcliente"`
	Nombrecli      string                      `json:"nombrecli"`
	Contribesp     bool                        `json:"contribesp"`
	RutaParme      bool                        `json:"ruta_parme"`
	Tipoprecio     int                         `json:"tipoprecio"`
	Emision        time.Time                   `json:"emision"`
	Recepcion      time.Time                   `json:"recepcion"`
	Vence          time.Time                   `json:"vence"`
	Diascred       int                         `json:"diascred"`
	Estatusdoc     int                         `json:"estatusdoc"`
	Dtotneto       float64                     `json:"dtotneto"`
	Dtotimpuest    float64                     `json:"dtotimpuest"`
	Dtotalfinal    float64                     `json:"dtotalfinal"`
	Dtotpagos      float64                     `json:"dtotpagos"`
	Dtotdescuen    float64                     `json:"dtotdescuen"`
	Dflete         float64                     `json:"dFlete"`
	Dtotdev        float64                     `json:"dtotdev"`
	Dvndmtototal   float64                     `json:"dvndmtototal"`
	Dretencion     float64                     `json:"dretencion"`
	Dretencioniva  float64                     `json:"dretencioniva"`
	Vendedor       string                      `json:"vendedor"`
	Codcoord       string                      `json:"codcoord"`
	Aceptadev      bool                        `json:"aceptadev"`
	KtiNegesp      bool                        `json:"kti_negesp"`
	Bsiva          float64                     `json:"bsiva"`
	Bsflete        float64                     `json:"bsflete"`
	Bsretencion    float64                     `json:"bsretencion"`
	Bsretencioniva float64                     `json:"bsretencioniva"`
	Tasadoc        float64                     `json:"tasadoc"`
	Mtodcto        float64                     `json:"mtodcto"`
	Fchvencedcto   time.Time                   `json:"fchvencedcto"`
	Tienedcto      bool                        `json:"tienedcto"`
	Cbsret         float64                     `json:"cbsret"`
	Cdret          float64                     `json:"cdret"`
	Cbsretiva      float64                     `json:"cbsretiva"`
	Cdretiva       float64                     `json:"cdretiva"`
	Cbsrparme      float64                     `json:"cbsrparme"`
	Cdrparme       float64                     `json:"cdrparme"`
	Cbsretflete    float64                     `json:"cbsretflete"`
	Cdretflete     float64                     `json:"cdretflete"`
	Bsmtoiva       float64                     `json:"bsmtoiva"`
	Bsmtofte       float64                     `json:"bsmtofte"`
	RetmunMto      float64                     `json:"retmun_mto"`
	Dolarflete     bool                        `json:"dolarflete"`
	Bsretflete     float64                     `json:"bsretflete"`
	Dretflete      float64                     `json:"dretflete"`
	DretmunMto     float64                     `json:"dretmun_mto"`
	Retivaoblig    bool                        `json:"retivaoblig"`
	Edoentrega     bool                        `json:"edoentrega"`
	Dmtoiva        float64                     `json:"dmtoiva"`
	Prcdctoaplic   float64                     `json:"prcdctoaplic"`
	Montodctodol   float64                     `json:"montodctodol"`
	Montodctobs    float64                     `json:"montodctobs"`
	Lines          []CreateDocumentLineRequest `json:"lines"`
}

type CreateDocumentLineRequest struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         string    `json:"grupo"`
	Subgrupo      string    `json:"subgrupo"`
	Origen        int       `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int       `json:"cantidad"`
	Cntdevuelt    int       `json:"cntdevuelt"`
	Vndcntdevuelt int       `json:"vndcntdevuelt"`
	Dvndmtototal  float64   `json:"dvndmtototal"`
	Dpreciofin    float64   `json:"dpreciofin"`
	Dpreciounit   float64   `json:"dpreciounit"`
	Dmontoneto    float64   `json:"dmontoneto"`
	Dmontototal   float64   `json:"dmontototal"`
	Timpueprc     float64   `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
}

type UpdateDocumentRequest struct {
	Agencia        string                      `json:"agencia"`
	Tipodoc        string                      `json:"tipodoc"`
	Documento      string                      `json:"documento"`
	Tipodocv       string                      `json:"tipodocv"`
	Codcliente     string                      `json:"codcliente"`
	Nombrecli      string                      `json:"nombrecli"`
	Contribesp     bool                        `json:"contribesp"`
	RutaParme      bool                        `json:"ruta_parme"`
	Tipoprecio     int                         `json:"tipoprecio"`
	Emision        time.Time                   `json:"emision"`
	Recepcion      time.Time                   `json:"recepcion"`
	Vence          time.Time                   `json:"vence"`
	Diascred       int                         `json:"diascred"`
	Estatusdoc     int                         `json:"estatusdoc"`
	Dtotneto       float64                     `json:"dtotneto"`
	Dtotimpuest    float64                     `json:"dtotimpuest"`
	Dtotalfinal    float64                     `json:"dtotalfinal"`
	Dtotpagos      float64                     `json:"dtotpagos"`
	Dtotdescuen    float64                     `json:"dtotdescuen"`
	Dflete         float64                     `json:"dFlete"`
	Dtotdev        float64                     `json:"dtotdev"`
	Dvndmtototal   float64                     `json:"dvndmtototal"`
	Dretencion     float64                     `json:"dretencion"`
	Dretencioniva  float64                     `json:"dretencioniva"`
	Vendedor       string                      `json:"vendedor"`
	Codcoord       string                      `json:"codcoord"`
	Aceptadev      bool                        `json:"aceptadev"`
	KtiNegesp      bool                        `json:"kti_negesp"`
	Bsiva          float64                     `json:"bsiva"`
	Bsflete        float64                     `json:"bsflete"`
	Bsretencion    float64                     `json:"bsretencion"`
	Bsretencioniva float64                     `json:"bsretencioniva"`
	Tasadoc        float64                     `json:"tasadoc"`
	Mtodcto        float64                     `json:"mtodcto"`
	Fchvencedcto   time.Time                   `json:"fchvencedcto"`
	Tienedcto      bool                        `json:"tienedcto"`
	Cbsret         float64                     `json:"cbsret"`
	Cdret          float64                     `json:"cdret"`
	Cbsretiva      float64                     `json:"cbsretiva"`
	Cdretiva       float64                     `json:"cdretiva"`
	Cbsrparme      float64                     `json:"cbsrparme"`
	Cdrparme       float64                     `json:"cdrparme"`
	Cbsretflete    float64                     `json:"cbsretflete"`
	Cdretflete     float64                     `json:"cdretflete"`
	Bsmtoiva       float64                     `json:"bsmtoiva"`
	Bsmtofte       float64                     `json:"bsmtofte"`
	RetmunMto      float64                     `json:"retmun_mto"`
	Dolarflete     bool                        `json:"dolarflete"`
	Bsretflete     float64                     `json:"bsretflete"`
	Dretflete      float64                     `json:"dretflete"`
	DretmunMto     float64                     `json:"dretmun_mto"`
	Retivaoblig    bool                        `json:"retivaoblig"`
	Edoentrega     bool                        `json:"edoentrega"`
	Dmtoiva        float64                     `json:"dmtoiva"`
	Prcdctoaplic   float64                     `json:"prcdctoaplic"`
	Montodctodol   float64                     `json:"montodctodol"`
	Montodctobs    float64                     `json:"montodctobs"`
	Lines          []UpdateDocumentLineRequest `json:"lines"`
}

type UpdateDocumentLineRequest struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         string    `json:"grupo"`
	Subgrupo      string    `json:"subgrupo"`
	Origen        int       `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int       `json:"cantidad"`
	Cntdevuelt    int       `json:"cntdevuelt"`
	Vndcntdevuelt int       `json:"vndcntdevuelt"`
	Dvndmtototal  float64   `json:"dvndmtototal"`
	Dpreciofin    float64   `json:"dpreciofin"`
	Dpreciounit   float64   `json:"dpreciounit"`
	Dmontoneto    float64   `json:"dmontoneto"`
	Dmontototal   float64   `json:"dmontototal"`
	Timpueprc     float64   `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
}

func DbDocToDocument(db *db.ClossDocument) *DocumentResponse {
	return &DocumentResponse{
		Agencia:        db.Agencia,
		Tipodoc:        db.Tipodoc,
		Documento:      db.Documento,
		Tipodocv:       db.Tipodocv,
		Codcliente:     db.Codcliente,
		Nombrecli:      db.Nombrecli,
		Contribesp:     db.Contribesp == 1,
		RutaParme:      db.RutaParme == 1,
		Tipoprecio:     int(db.Tipoprecio),
		Emision:        db.Emision,
		Recepcion:      db.Recepcion,
		Vence:          db.Vence,
		Diascred:       int(db.Diascred),
		Estatusdoc:     int(db.Estatusdoc),
		Dtotneto:       db.Dtotneto,
		Dtotimpuest:    db.Dtotimpuest,
		Dtotalfinal:    db.Dtotalfinal,
		Dtotpagos:      db.Dtotpagos,
		Dtotdescuen:    db.Dtotdescuen,
		Dflete:         db.Dflete,
		Dtotdev:        db.Dtotdev,
		Dvndmtototal:   db.Dvndmtototal,
		Dretencion:     db.Dretencion,
		Dretencioniva:  db.Dretencioniva,
		Vendedor:       db.Vendedor,
		Codcoord:       db.Codcoord,
		Aceptadev:      db.Aceptadev == 1,
		KtiNegesp:      db.KtiNegesp == 1,
		Bsiva:          db.Bsiva,
		Bsflete:        db.Bsflete,
		Bsretencion:    db.Bsretencion,
		Bsretencioniva: db.Bsretencioniva,
		Tasadoc:        db.Tasadoc,
		Mtodcto:        db.Mtodcto,
		Fchvencedcto:   db.Fchvencedcto,
		Tienedcto:      db.Tienedcto == 1,
		Cbsret:         db.Cbsret,
		Cdret:          db.Cdret,
		Cbsretiva:      db.Cbsretiva,
		Cdretiva:       db.Cdretiva,
		Cbsrparme:      db.Cbsrparme,
		Cdrparme:       db.Cdrparme,
		Cbsretflete:    db.Cbsretflete,
		Cdretflete:     db.Cdretflete,
		Bsmtoiva:       db.Bsmtoiva,
		Bsmtofte:       db.Bsmtofte,
		RetmunMto:      db.RetmunMto,
		Dolarflete:     db.Dolarflete == 1,
		Bsretflete:     db.Bsretflete,
		Dretflete:      db.Dretflete,
		DretmunMto:     db.DretmunMto,
		Retivaoblig:    db.Retivaoblig == 1,
		Edoentrega:     db.Edoentrega == 1,
		Dmtoiva:        db.Dmtoiva,
		Prcdctoaplic:   db.Prcdctoaplic,
		Montodctodol:   db.Montodctodol,
		Montodctobs:    db.Montodctobs,
		CreatedAt:      db.CreatedAt,
		UpdatedAt:      db.UpdatedAt,
	}
}

func DbDocLineToDocLine(db *db.ClossDocumentLine) *DocumentLineResponse {
	return &DocumentLineResponse{
		Agencia:       db.Agencia,
		Tipodoc:       db.Tipodoc,
		Documento:     db.Documento,
		Tipodocv:      db.Tipodocv,
		Grupo:         db.Grupo,
		Subgrupo:      db.Subgrupo,
		Origen:        int(db.Origen),
		Codigo:        db.Codigo,
		Codhijo:       db.Codhijo,
		Pid:           db.Pid,
		Nombre:        db.Nombre,
		Cantidad:      int(db.Cantidad),
		Cntdevuelt:    int(db.Cntdevuelt),
		Vndcntdevuelt: int(db.Vndcntdevuelt),
		Dvndmtototal:  db.Dvndmtototal,
		Dpreciofin:    db.Dpreciofin,
		Dpreciounit:   db.Dpreciounit,
		Dmontoneto:    db.Dmontoneto,
		Dmontototal:   db.Dmontototal,
		Timpueprc:     db.Timpueprc,
		Unidevuelt:    int(db.Unidevuelt),
		Fechadoc:      db.Fechadoc,
		Vendedor:      db.Vendedor,
		Codcoord:      db.Codcoord,
	}
}

func CreateDocumentToDb(r *CreateDocumentRequest) *db.CreateDocumentParams {
	return &db.CreateDocumentParams{
		Agencia:        r.Agencia,
		Tipodoc:        r.Tipodoc,
		Documento:      r.Documento,
		Tipodocv:       r.Tipodocv,
		Codcliente:     r.Codcliente,
		Nombrecli:      r.Nombrecli,
		Contribesp:     int64(util.BoolToInt(r.Contribesp)),
		RutaParme:      int64(util.BoolToInt(r.RutaParme)),
		Tipoprecio:     int64(r.Tipoprecio),
		Emision:        r.Emision.String(),
		Recepcion:      r.Recepcion.String(),
		Vence:          r.Vence.String(),
		Diascred:       int64(r.Diascred),
		Estatusdoc:     int64(r.Estatusdoc),
		Dtotneto:       r.Dtotneto,
		Dtotimpuest:    r.Dtotimpuest,
		Dtotalfinal:    r.Dtotalfinal,
		Dtotpagos:      r.Dtotpagos,
		Dtotdescuen:    r.Dtotdescuen,
		Dflete:         r.Dflete,
		Dtotdev:        r.Dtotdev,
		Dvndmtototal:   r.Dvndmtototal,
		Dretencion:     r.Dretencion,
		Dretencioniva:  r.Dretencioniva,
		Vendedor:       r.Vendedor,
		Codcoord:       r.Codcoord,
		Aceptadev:      int64(util.BoolToInt(r.Aceptadev)),
		KtiNegesp:      int64(util.BoolToInt(r.KtiNegesp)),
		Bsiva:          r.Bsiva,
		Bsflete:        r.Bsflete,
		Bsretencion:    r.Bsretencion,
		Bsretencioniva: r.Bsretencioniva,
		Tasadoc:        r.Tasadoc,
		Mtodcto:        r.Mtodcto,
		Fchvencedcto:   r.Fchvencedcto.String(),
		Tienedcto:      int64(util.BoolToInt(r.Tienedcto)),
		Cbsret:         r.Cbsret,
		Cdret:          r.Cdret,
		Cbsretiva:      r.Cbsretiva,
		Cdretiva:       r.Cdretiva,
		Cbsrparme:      r.Cbsrparme,
		Cdrparme:       r.Cdrparme,
		Cbsretflete:    r.Cbsretflete,
		Cdretflete:     r.Cdretflete,
		Bsmtoiva:       r.Bsmtoiva,
		Bsmtofte:       r.Bsmtofte,
		RetmunMto:      r.RetmunMto,
		Dolarflete:     int64(util.BoolToInt(r.Dolarflete)),
		Bsretflete:     r.Bsretflete,
		Dretflete:      r.Dretflete,
		DretmunMto:     r.DretmunMto,
		Retivaoblig:    int64(util.BoolToInt(r.Retivaoblig)),
		Edoentrega:     int64(util.BoolToInt(r.Edoentrega)),
		Dmtoiva:        r.Dmtoiva,
		Prcdctoaplic:   r.Prcdctoaplic,
		Montodctodol:   r.Montodctodol,
		Montodctobs:    r.Montodctobs,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      time.Now().String(),
	}
}

func UpdateDocumentToDb(r *UpdateDocumentRequest) *db.UpdateDocumentParams {
	return &db.UpdateDocumentParams{
		Agencia:        r.Agencia,
		Tipodoc:        r.Tipodoc,
		Documento:      r.Documento,
		Tipodocv:       r.Tipodocv,
		Codcliente:     r.Codcliente,
		Nombrecli:      r.Nombrecli,
		Contribesp:     int64(util.BoolToInt(r.Contribesp)),
		RutaParme:      int64(util.BoolToInt(r.RutaParme)),
		Tipoprecio:     int64(r.Tipoprecio),
		Emision:        r.Emision.String(),
		Recepcion:      r.Recepcion.String(),
		Vence:          r.Vence.String(),
		Diascred:       int64(r.Diascred),
		Estatusdoc:     int64(r.Estatusdoc),
		Dtotneto:       r.Dtotneto,
		Dtotimpuest:    r.Dtotimpuest,
		Dtotalfinal:    r.Dtotalfinal,
		Dtotpagos:      r.Dtotpagos,
		Dtotdescuen:    r.Dtotdescuen,
		Dflete:         r.Dflete,
		Dtotdev:        r.Dtotdev,
		Dvndmtototal:   r.Dvndmtototal,
		Dretencion:     r.Dretencion,
		Dretencioniva:  r.Dretencioniva,
		Vendedor:       r.Vendedor,
		Codcoord:       r.Codcoord,
		Aceptadev:      int64(util.BoolToInt(r.Aceptadev)),
		KtiNegesp:      int64(util.BoolToInt(r.KtiNegesp)),
		Bsiva:          r.Bsiva,
		Bsflete:        r.Bsflete,
		Bsretencion:    r.Bsretencion,
		Bsretencioniva: r.Bsretencioniva,
		Tasadoc:        r.Tasadoc,
		Mtodcto:        r.Mtodcto,
		Fchvencedcto:   r.Fchvencedcto.String(),
		Tienedcto:      int64(util.BoolToInt(r.Tienedcto)),
		Cbsret:         r.Cbsret,
		Cdret:          r.Cdret,
		Cbsretiva:      r.Cbsretiva,
		Cdretiva:       r.Cdretiva,
		Cbsrparme:      r.Cbsrparme,
		Cdrparme:       r.Cdrparme,
		Cbsretflete:    r.Cbsretflete,
		Cdretflete:     r.Cdretflete,
		Bsmtoiva:       r.Bsmtoiva,
		Bsmtofte:       r.Bsmtofte,
		RetmunMto:      r.RetmunMto,
		Dolarflete:     int64(util.BoolToInt(r.Dolarflete)),
		Bsretflete:     r.Bsretflete,
		Dretflete:      r.Dretflete,
		DretmunMto:     r.DretmunMto,
		Retivaoblig:    int64(util.BoolToInt(r.Retivaoblig)),
		Edoentrega:     int64(util.BoolToInt(r.Edoentrega)),
		Dmtoiva:        r.Dmtoiva,
		Prcdctoaplic:   r.Prcdctoaplic,
		Montodctodol:   r.Montodctodol,
		Montodctobs:    r.Montodctobs,
		UpdatedAt:      time.Now().String(),
	}
}

func CreateDocumentLinesToDb(r *CreateDocumentLineRequest) *db.CreateDocumentLineParams {
	return &db.CreateDocumentLineParams{
		Agencia:       r.Agencia,
		Tipodoc:       r.Tipodoc,
		Documento:     r.Documento,
		Tipodocv:      r.Tipodocv,
		Grupo:         r.Grupo,
		Subgrupo:      r.Subgrupo,
		Origen:        int64(r.Origen),
		Codigo:        r.Codigo,
		Codhijo:       r.Codhijo,
		Pid:           r.Pid,
		Nombre:        r.Nombre,
		Cantidad:      int64(r.Cantidad),
		Cntdevuelt:    int64(r.Cntdevuelt),
		Vndcntdevuelt: int64(r.Vndcntdevuelt),
		Dvndmtototal:  r.Dvndmtototal,
		Dpreciofin:    r.Dpreciofin,
		Dpreciounit:   r.Dpreciounit,
		Dmontoneto:    r.Dmontoneto,
		Dmontototal:   r.Dmontototal,
		Timpueprc:     r.Timpueprc,
		Unidevuelt:    int64(r.Unidevuelt),
		Fechadoc:      r.Fechadoc.String(),
		Vendedor:      r.Vendedor,
		Codcoord:      r.Codcoord,
	}
}

func UpdateDocumentLinesToDb(r *UpdateDocumentLineRequest) *db.UpdateDocumentLineParams {
	return &db.UpdateDocumentLineParams{
		Agencia:       r.Agencia,
		Tipodoc:       r.Tipodoc,
		Documento:     r.Documento,
		Tipodocv:      r.Tipodocv,
		Grupo:         r.Grupo,
		Subgrupo:      r.Subgrupo,
		Origen:        int64(r.Origen),
		Codigo:        r.Codigo,
		Codhijo:       r.Codhijo,
		Pid:           r.Pid,
		Nombre:        r.Nombre,
		Cantidad:      int64(r.Cantidad),
		Cntdevuelt:    int64(r.Cntdevuelt),
		Vndcntdevuelt: int64(r.Vndcntdevuelt),
		Dvndmtototal:  r.Dvndmtototal,
		Dpreciofin:    r.Dpreciofin,
		Dpreciounit:   r.Dpreciounit,
		Dmontoneto:    r.Dmontoneto,
		Dmontototal:   r.Dmontototal,
		Timpueprc:     r.Timpueprc,
		Unidevuelt:    int64(r.Unidevuelt),
		Fechadoc:      r.Fechadoc.String(),
		Vendedor:      r.Vendedor,
		Codcoord:      r.Codcoord,
	}
}

func DocToDocWithLines(key *DocumentResponse, value []DocumentLineResponse) *DocumentWithLinesResponse {
	return &DocumentWithLinesResponse{
		DocumentResponse: *key,
		Lines:            value,
	}
}
