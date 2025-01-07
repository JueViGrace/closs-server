package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type DocumentWithLines struct {
	Document
	Lines []DocumentLine `json:"lines"`
}

type Document struct {
	Agencia        string    `json:"agencia"`
	Tipodoc        string    `json:"tipodoc"`
	Documento      string    `json:"documento"`
	Tipodocv       string    `json:"tipodocv"`
	Codcliente     string    `json:"codcliente"`
	Nombrecli      string    `json:"nombrecli"`
	Contribesp     bool      `json:"contribesp"`
	RutaParme      bool      `json:"ruta_parme"`
	Tipoprecio     string    `json:"tipoprecio"`
	Emision        time.Time `json:"emision"`
	Recepcion      time.Time `json:"recepcion"`
	Vence          time.Time `json:"vence"`
	Diascred       string    `json:"diascred"`
	Estatusdoc     string    `json:"estatusdoc"`
	Dtotneto       string    `json:"dtotneto"`
	Dtotimpuest    string    `json:"dtotimpuest"`
	Dtotalfinal    string    `json:"dtotalfinal"`
	Dtotpagos      string    `json:"dtotpagos"`
	Dtotdescuen    string    `json:"dtotdescuen"`
	Dflete         string    `json:"dFlete"`
	Dtotdev        string    `json:"dtotdev"`
	Dvndmtototal   string    `json:"dvndmtototal"`
	Dretencion     string    `json:"dretencion"`
	Dretencioniva  string    `json:"dretencioniva"`
	Vendedor       string    `json:"vendedor"`
	Codcoord       string    `json:"codcoord"`
	Aceptadev      bool      `json:"aceptadev"`
	KtiNegesp      bool      `json:"kti_negesp"`
	Bsiva          string    `json:"bsiva"`
	Bsflete        string    `json:"bsflete"`
	Bsretencion    string    `json:"bsretencion"`
	Bsretencioniva string    `json:"bsretencioniva"`
	Tasadoc        string    `json:"tasadoc"`
	Mtodcto        string    `json:"mtodcto"`
	Fchvencedcto   time.Time `json:"fchvencedcto"`
	Tienedcto      bool      `json:"tienedcto"`
	Cbsret         string    `json:"cbsret"`
	Cdret          string    `json:"cdret"`
	Cbsretiva      string    `json:"cbsretiva"`
	Cdretiva       string    `json:"cdretiva"`
	Cbsrparme      string    `json:"cbsrparme"`
	Cdrparme       string    `json:"cdrparme"`
	Cbsretflete    string    `json:"cbsretflete"`
	Cdretflete     string    `json:"cdretflete"`
	Bsmtoiva       string    `json:"bsmtoiva"`
	Bsmtofte       string    `json:"bsmtofte"`
	RetmunMto      string    `json:"retmun_mto"`
	Dolarflete     bool      `json:"dolarflete"`
	Bsretflete     string    `json:"bsretflete"`
	Dretflete      string    `json:"dretflete"`
	DretmunMto     string    `json:"dretmun_mto"`
	Retivaoblig    bool      `json:"retivaoblig"`
	Edoentrega     bool      `json:"edoentrega"`
	Dmtoiva        string    `json:"dmtoiva"`
	Prcdctoaplic   string    `json:"prcdctoaplic"`
	Montodctodol   string    `json:"montodctodol"`
	Montodctobs    string    `json:"montodctobs"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DocumentLine struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         int32     `json:"grupo"`
	Subgrupo      int32     `json:"subgrupo"`
	Origen        string    `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int32     `json:"cantidad"`
	Cntdevuelt    int32     `json:"cntdevuelt"`
	Vndcntdevuelt string    `json:"vndcntdevuelt"`
	Dvndmtototal  string    `json:"dvndmtototal"`
	Dpreciofin    string    `json:"dpreciofin"`
	Dpreciounit   string    `json:"dpreciounit"`
	Dmontoneto    string    `json:"dmontoneto"`
	Dmontototal   string    `json:"dmontototal"`
	Timpueprc     string    `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateDocumentRequest struct {
	Agencia        string    `json:"agencia"`
	Tipodoc        string    `json:"tipodoc"`
	Documento      string    `json:"documento"`
	Tipodocv       string    `json:"tipodocv"`
	Codcliente     string    `json:"codcliente"`
	Nombrecli      string    `json:"nombrecli"`
	Contribesp     bool      `json:"contribesp"`
	RutaParme      bool      `json:"ruta_parme"`
	Tipoprecio     string    `json:"tipoprecio"`
	Emision        time.Time `json:"emision"`
	Recepcion      time.Time `json:"recepcion"`
	Vence          time.Time `json:"vence"`
	Diascred       string    `json:"diascred"`
	Estatusdoc     string    `json:"estatusdoc"`
	Dtotneto       string    `json:"dtotneto"`
	Dtotimpuest    string    `json:"dtotimpuest"`
	Dtotalfinal    string    `json:"dtotalfinal"`
	Dtotpagos      string    `json:"dtotpagos"`
	Dtotdescuen    string    `json:"dtotdescuen"`
	Dflete         string    `json:"dFlete"`
	Dtotdev        string    `json:"dtotdev"`
	Dvndmtototal   string    `json:"dvndmtototal"`
	Dretencion     string    `json:"dretencion"`
	Dretencioniva  string    `json:"dretencioniva"`
	Vendedor       string    `json:"vendedor"`
	Codcoord       string    `json:"codcoord"`
	Aceptadev      bool      `json:"aceptadev"`
	KtiNegesp      bool      `json:"kti_negesp"`
	Bsiva          string    `json:"bsiva"`
	Bsflete        string    `json:"bsflete"`
	Bsretencion    string    `json:"bsretencion"`
	Bsretencioniva string    `json:"bsretencioniva"`
	Tasadoc        string    `json:"tasadoc"`
	Mtodcto        string    `json:"mtodcto"`
	Fchvencedcto   time.Time `json:"fchvencedcto"`
	Tienedcto      bool      `json:"tienedcto"`
	Cbsret         string    `json:"cbsret"`
	Cdret          string    `json:"cdret"`
	Cbsretiva      string    `json:"cbsretiva"`
	Cdretiva       string    `json:"cdretiva"`
	Cbsrparme      string    `json:"cbsrparme"`
	Cdrparme       string    `json:"cdrparme"`
	Cbsretflete    string    `json:"cbsretflete"`
	Cdretflete     string    `json:"cdretflete"`
	Bsmtoiva       string    `json:"bsmtoiva"`
	Bsmtofte       string    `json:"bsmtofte"`
	RetmunMto      string    `json:"retmun_mto"`
	Dolarflete     bool      `json:"dolarflete"`
	Bsretflete     string    `json:"bsretflete"`
	Dretflete      string    `json:"dretflete"`
	DretmunMto     string    `json:"dretmun_mto"`
	Retivaoblig    bool      `json:"retivaoblig"`
	Edoentrega     bool      `json:"edoentrega"`
	Dmtoiva        string    `json:"dmtoiva"`
	Prcdctoaplic   string    `json:"prcdctoaplic"`
	Montodctodol   string    `json:"montodctodol"`
	Montodctobs    string    `json:"montodctobs"`
}

type CreateDocumentLineRequest struct {
	DocID         uuid.UUID `json:"doc_id"`
	ArticuloID    uuid.UUID `json:"articulo_id"`
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         int32     `json:"grupo"`
	Subgrupo      int32     `json:"subgrupo"`
	Origen        string    `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      int32     `json:"cantidad"`
	Cntdevuelt    int32     `json:"cntdevuelt"`
	Vndcntdevuelt string    `json:"vndcntdevuelt"`
	Dvndmtototal  string    `json:"dvndmtototal"`
	Dpreciofin    string    `json:"dpreciofin"`
	Dpreciounit   string    `json:"dpreciounit"`
	Dmontoneto    string    `json:"dmontoneto"`
	Dmontototal   string    `json:"dmontototal"`
	Timpueprc     string    `json:"timpueprc"`
	Unidevuelt    int32     `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
}

type UpdateDocumentRequest struct{}

type UpdateDocumentLineRequest struct{}

func DbKeDoccToDocument(db *db.ClossDocument) *Document {
	return &Document{
		Agencia:        db.Agencia,
		Tipodoc:        db.Tipodoc,
		Documento:      db.Documento,
		Tipodocv:       db.Tipodocv,
		Codcliente:     db.Codcliente,
		Nombrecli:      db.Nombrecli,
		Contribesp:     db.Contribesp == 1,
		RutaParme:      db.RutaParme == 1,
		Tipoprecio:     db.Tipoprecio,
		Emision:        db.Emision,
		Recepcion:      db.Recepcion,
		Vence:          db.Vence,
		Diascred:       db.Diascred,
		Estatusdoc:     db.Estatusdoc,
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
		Aceptadev:      db.Aceptadev,
		KtiNegesp:      db.KtiNegesp,
		Bsiva:          db.Bsiva,
		Bsflete:        db.Bsflete,
		Bsretencion:    db.Bsretencion,
		Bsretencioniva: db.Bsretencioniva,
		Tasadoc:        db.Tasadoc,
		Mtodcto:        db.Mtodcto,
		Fchvencedcto:   db.Fchvencedcto,
		Tienedcto:      db.Tienedcto,
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
		Dolarflete:     db.Dolarflete,
		Bsretflete:     db.Bsretflete,
		Dretflete:      db.Dretflete,
		DretmunMto:     db.DretmunMto,
		Retivaoblig:    db.Retivaoblig,
		Edoentrega:     db.Edoentrega,
		Dmtoiva:        db.Dmtoiva,
		Prcdctoaplic:   db.Prcdctoaplic,
		Montodctodol:   db.Montodctodol,
		Montodctobs:    db.Montodctobs,
		CreatedAt:      db.CreatedAt,
		UpdatedAt:      db.UpdatedAt,
	}
}

func DocMapToDocWithLines(key *Document, value *[]DocumentLine) *DocumentWithLines {
	return &DocumentWithLines{
		Document: Document{
			Agencia:        key.Agencia,
			Tipodoc:        key.Tipodoc,
			Documento:      key.Documento,
			Tipodocv:       key.Tipodocv,
			Codcliente:     key.Codcliente,
			Nombrecli:      key.Nombrecli,
			Contribesp:     key.Contribesp,
			RutaParme:      key.RutaParme,
			Tipoprecio:     key.Tipoprecio,
			Emision:        key.Emision,
			Recepcion:      key.Recepcion,
			Vence:          key.Vence,
			Diascred:       key.Diascred,
			Estatusdoc:     key.Estatusdoc,
			Dtotneto:       key.Dtotneto,
			Dtotimpuest:    key.Dtotimpuest,
			Dtotalfinal:    key.Dtotalfinal,
			Dtotpagos:      key.Dtotpagos,
			Dtotdescuen:    key.Dtotdescuen,
			Dflete:         key.Dflete,
			Dtotdev:        key.Dtotdev,
			Dvndmtototal:   key.Dvndmtototal,
			Dretencion:     key.Dretencion,
			Dretencioniva:  key.Dretencioniva,
			Vendedor:       key.Vendedor,
			Codcoord:       key.Codcoord,
			Aceptadev:      key.Aceptadev,
			KtiNegesp:      key.KtiNegesp,
			Bsiva:          key.Bsiva,
			Bsflete:        key.Bsflete,
			Bsretencion:    key.Bsretencion,
			Bsretencioniva: key.Bsretencioniva,
			Tasadoc:        key.Tasadoc,
			Mtodcto:        key.Mtodcto,
			Fchvencedcto:   key.Fchvencedcto,
			Tienedcto:      key.Tienedcto,
			Cbsret:         key.Cbsret,
			Cdret:          key.Cdret,
			Cbsretiva:      key.Cbsretiva,
			Cdretiva:       key.Cdretiva,
			Cbsrparme:      key.Cbsrparme,
			Cdrparme:       key.Cdrparme,
			Cbsretflete:    key.Cbsretflete,
			Cdretflete:     key.Cdretflete,
			Bsmtoiva:       key.Bsmtoiva,
			Bsmtofte:       key.Bsmtofte,
			RetmunMto:      key.RetmunMto,
			Dolarflete:     key.Dolarflete,
			Bsretflete:     key.Bsretflete,
			Dretflete:      key.Dretflete,
			DretmunMto:     key.DretmunMto,
			Retivaoblig:    key.Retivaoblig,
			Edoentrega:     key.Edoentrega,
			Dmtoiva:        key.Dmtoiva,
			Prcdctoaplic:   key.Prcdctoaplic,
			Montodctodol:   key.Montodctodol,
			Montodctobs:    key.Montodctobs,
			CreatedAt:      key.CreatedAt,
			UpdatedAt:      key.UpdatedAt,
			DeletedAt:      key.DeletedAt,
		},
		Lines: *value,
	}
}

func DbDocToDocument(db *db.AdminGetDocumentsWithLinesRow) *Document {
	return &Document{
		Agencia:        db.Agencia,
		Tipodoc:        db.Tipodoc,
		Documento:      db.Documento,
		Tipodocv:       db.Tipodocv,
		Codcliente:     db.Codcliente,
		Nombrecli:      db.Nombrecli,
		Contribesp:     db.Contribesp,
		RutaParme:      db.RutaParme,
		Tipoprecio:     db.Tipoprecio,
		Emision:        db.Emision,
		Recepcion:      db.Recepcion,
		Vence:          db.Vence,
		Diascred:       db.Diascred,
		Estatusdoc:     db.Estatusdoc,
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
		Aceptadev:      db.Aceptadev,
		KtiNegesp:      db.KtiNegesp,
		Bsiva:          db.Bsiva,
		Bsflete:        db.Bsflete,
		Bsretencion:    db.Bsretencion,
		Bsretencioniva: db.Bsretencioniva,
		Tasadoc:        db.Tasadoc,
		Mtodcto:        db.Mtodcto,
		Fchvencedcto:   db.Fchvencedcto,
		Tienedcto:      db.Tienedcto,
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
		Dolarflete:     db.Dolarflete,
		Bsretflete:     db.Bsretflete,
		Dretflete:      db.Dretflete,
		DretmunMto:     db.DretmunMto,
		Retivaoblig:    db.Retivaoblig,
		Edoentrega:     db.Edoentrega,
		Dmtoiva:        db.Dmtoiva,
		Prcdctoaplic:   db.Prcdctoaplic,
		Montodctodol:   db.Montodctodol,
		Montodctobs:    db.Montodctobs,
		CreatedAt:      db.CreatedAt,
		UpdatedAt:      db.UpdatedAt,
		DeletedAt:      db.DeletedAt.Time,
	}
}

func DbDocToDocLine(db *db.AdminGetDocumentsWithLinesRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       db.Agencia_2.String,
		Tipodoc:       db.Tipodoc_2.String,
		Documento:     db.Documento_2.String,
		Tipodocv:      db.Tipodocv_2.String,
		Grupo:         db.Grupo.Int32,
		Subgrupo:      db.Subgrupo.Int32,
		Origen:        db.Origen.String,
		Codigo:        db.Codigo.String,
		Codhijo:       db.Codhijo.String,
		Pid:           db.Pid.String,
		Nombre:        db.Nombre.String,
		Cantidad:      db.Cantidad.Int32,
		Cntdevuelt:    db.Cntdevuelt.Int32,
		Vndcntdevuelt: db.Vndcntdevuelt.String,
		Dvndmtototal:  db.Dvndmtototal_2.String,
		Dpreciofin:    db.Dpreciofin.String,
		Dpreciounit:   db.Dpreciounit.String,
		Dmontoneto:    db.Dmontoneto.String,
		Dmontototal:   db.Dmontototal.String,
		Timpueprc:     db.Timpueprc.String,
		Unidevuelt:    db.Unidevuelt.Int32,
		Fechadoc:      db.Fechadoc.Time,
		Vendedor:      db.Vendedor_2.String,
		Codcoord:      db.Codcoord_2.String,
		CreatedAt:     db.CreatedAt_2.Time,
		UpdatedAt:     db.UpdatedAt_2.Time,
		DeletedAt:     db.DeletedAt_2.Time,
	}
}

func DbDocByCodeToDocument(db *db.GetDocumentsWithLinesBySalesmanRow) *Document {
	return &Document{
		Agencia:        db.Agencia,
		Tipodoc:        db.Tipodoc,
		Documento:      db.Documento,
		Tipodocv:       db.Tipodocv,
		Codcliente:     db.Codcliente,
		Nombrecli:      db.Nombrecli,
		Contribesp:     db.Contribesp,
		RutaParme:      db.RutaParme,
		Tipoprecio:     db.Tipoprecio,
		Emision:        db.Emision,
		Recepcion:      db.Recepcion,
		Vence:          db.Vence,
		Diascred:       db.Diascred,
		Estatusdoc:     db.Estatusdoc,
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
		Aceptadev:      db.Aceptadev,
		KtiNegesp:      db.KtiNegesp,
		Bsiva:          db.Bsiva,
		Bsflete:        db.Bsflete,
		Bsretencion:    db.Bsretencion,
		Bsretencioniva: db.Bsretencioniva,
		Tasadoc:        db.Tasadoc,
		Mtodcto:        db.Mtodcto,
		Fchvencedcto:   db.Fchvencedcto,
		Tienedcto:      db.Tienedcto,
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
		Dolarflete:     db.Dolarflete,
		Bsretflete:     db.Bsretflete,
		Dretflete:      db.Dretflete,
		DretmunMto:     db.DretmunMto,
		Retivaoblig:    db.Retivaoblig,
		Edoentrega:     db.Edoentrega,
		Dmtoiva:        db.Dmtoiva,
		Prcdctoaplic:   db.Prcdctoaplic,
		Montodctodol:   db.Montodctodol,
		Montodctobs:    db.Montodctobs,
		CreatedAt:      db.CreatedAt,
		UpdatedAt:      db.UpdatedAt,
		DeletedAt:      db.DeletedAt.Time,
	}
}

func DbDocByCodeToDocLine(db *db.GetDocumentsWithLinesBySalesmanRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       db.Agencia_2.String,
		Tipodoc:       db.Tipodoc_2.String,
		Documento:     db.Documento_2.String,
		Tipodocv:      db.Tipodocv_2.String,
		Grupo:         db.Grupo.Int32,
		Subgrupo:      db.Subgrupo.Int32,
		Origen:        db.Origen.String,
		Codigo:        db.Codigo.String,
		Codhijo:       db.Codhijo.String,
		Pid:           db.Pid.String,
		Nombre:        db.Nombre.String,
		Cantidad:      db.Cantidad.Int32,
		Cntdevuelt:    db.Cntdevuelt.Int32,
		Vndcntdevuelt: db.Vndcntdevuelt.String,
		Dvndmtototal:  db.Dvndmtototal_2.String,
		Dpreciofin:    db.Dpreciofin.String,
		Dpreciounit:   db.Dpreciounit.String,
		Dmontoneto:    db.Dmontoneto.String,
		Dmontototal:   db.Dmontototal.String,
		Timpueprc:     db.Timpueprc.String,
		Unidevuelt:    db.Unidevuelt.Int32,
		Fechadoc:      db.Fechadoc.Time,
		Vendedor:      db.Vendedor_2.String,
		Codcoord:      db.Codcoord_2.String,
		CreatedAt:     db.CreatedAt,
		UpdatedAt:     db.UpdatedAt,
		DeletedAt:     db.DeletedAt.Time,
	}
}
