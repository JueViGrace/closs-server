package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
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
	DeletedAt      time.Time `json:"-"`
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
	DeletedAt     time.Time `json:"-"`
}

func DbKeDoccToDocument(kd *db.KeDoccti) *Document {
	return &Document{
		Agencia:        kd.Agencia,
		Tipodoc:        kd.Tipodoc,
		Documento:      kd.Documento,
		Tipodocv:       kd.Tipodocv,
		Codcliente:     kd.Codcliente,
		Nombrecli:      kd.Nombrecli,
		Contribesp:     kd.Contribesp,
		RutaParme:      kd.RutaParme,
		Tipoprecio:     kd.Tipoprecio,
		Emision:        kd.Emision,
		Recepcion:      kd.Recepcion,
		Vence:          kd.Vence,
		Diascred:       kd.Diascred,
		Estatusdoc:     kd.Estatusdoc,
		Dtotneto:       kd.Dtotneto,
		Dtotimpuest:    kd.Dtotimpuest,
		Dtotalfinal:    kd.Dtotalfinal,
		Dtotpagos:      kd.Dtotpagos,
		Dtotdescuen:    kd.Dtotdescuen,
		Dflete:         kd.Dflete,
		Dtotdev:        kd.Dtotdev,
		Dvndmtototal:   kd.Dvndmtototal,
		Dretencion:     kd.Dretencion,
		Dretencioniva:  kd.Dretencioniva,
		Vendedor:       kd.Vendedor,
		Codcoord:       kd.Codcoord,
		Aceptadev:      kd.Aceptadev,
		KtiNegesp:      kd.KtiNegesp,
		Bsiva:          kd.Bsiva,
		Bsflete:        kd.Bsflete,
		Bsretencion:    kd.Bsretencion,
		Bsretencioniva: kd.Bsretencioniva,
		Tasadoc:        kd.Tasadoc,
		Mtodcto:        kd.Mtodcto,
		Fchvencedcto:   kd.Fchvencedcto,
		Tienedcto:      kd.Tienedcto,
		Cbsret:         kd.Cbsret,
		Cdret:          kd.Cdret,
		Cbsretiva:      kd.Cbsretiva,
		Cdretiva:       kd.Cdretiva,
		Cbsrparme:      kd.Cbsrparme,
		Cdrparme:       kd.Cdrparme,
		Cbsretflete:    kd.Cbsretflete,
		Cdretflete:     kd.Cdretflete,
		Bsmtoiva:       kd.Bsmtoiva,
		Bsmtofte:       kd.Bsmtofte,
		RetmunMto:      kd.RetmunMto,
		Dolarflete:     kd.Dolarflete,
		Bsretflete:     kd.Bsretflete,
		Dretflete:      kd.Dretflete,
		DretmunMto:     kd.DretmunMto,
		Retivaoblig:    kd.Retivaoblig,
		Edoentrega:     kd.Edoentrega,
		Dmtoiva:        kd.Dmtoiva,
		Prcdctoaplic:   kd.Prcdctoaplic,
		Montodctodol:   kd.Montodctodol,
		Montodctobs:    kd.Montodctobs,
		CreatedAt:      kd.CreatedAt,
		UpdatedAt:      kd.UpdatedAt,
		DeletedAt:      kd.DeletedAt.Time,
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

func DbDocToDocument(dbDoc *db.AdminGetDocumentsWithLinesRow) *Document {
	return &Document{
		Agencia:        dbDoc.Agencia,
		Tipodoc:        dbDoc.Tipodoc,
		Documento:      dbDoc.Documento,
		Tipodocv:       dbDoc.Tipodocv,
		Codcliente:     dbDoc.Codcliente,
		Nombrecli:      dbDoc.Nombrecli,
		Contribesp:     dbDoc.Contribesp,
		RutaParme:      dbDoc.RutaParme,
		Tipoprecio:     dbDoc.Tipoprecio,
		Emision:        dbDoc.Emision,
		Recepcion:      dbDoc.Recepcion,
		Vence:          dbDoc.Vence,
		Diascred:       dbDoc.Diascred,
		Estatusdoc:     dbDoc.Estatusdoc,
		Dtotneto:       dbDoc.Dtotneto,
		Dtotimpuest:    dbDoc.Dtotimpuest,
		Dtotalfinal:    dbDoc.Dtotalfinal,
		Dtotpagos:      dbDoc.Dtotpagos,
		Dtotdescuen:    dbDoc.Dtotdescuen,
		Dflete:         dbDoc.Dflete,
		Dtotdev:        dbDoc.Dtotdev,
		Dvndmtototal:   dbDoc.Dvndmtototal,
		Dretencion:     dbDoc.Dretencion,
		Dretencioniva:  dbDoc.Dretencioniva,
		Vendedor:       dbDoc.Vendedor,
		Codcoord:       dbDoc.Codcoord,
		Aceptadev:      dbDoc.Aceptadev,
		KtiNegesp:      dbDoc.KtiNegesp,
		Bsiva:          dbDoc.Bsiva,
		Bsflete:        dbDoc.Bsflete,
		Bsretencion:    dbDoc.Bsretencion,
		Bsretencioniva: dbDoc.Bsretencioniva,
		Tasadoc:        dbDoc.Tasadoc,
		Mtodcto:        dbDoc.Mtodcto,
		Fchvencedcto:   dbDoc.Fchvencedcto,
		Tienedcto:      dbDoc.Tienedcto,
		Cbsret:         dbDoc.Cbsret,
		Cdret:          dbDoc.Cdret,
		Cbsretiva:      dbDoc.Cbsretiva,
		Cdretiva:       dbDoc.Cdretiva,
		Cbsrparme:      dbDoc.Cbsrparme,
		Cdrparme:       dbDoc.Cdrparme,
		Cbsretflete:    dbDoc.Cbsretflete,
		Cdretflete:     dbDoc.Cdretflete,
		Bsmtoiva:       dbDoc.Bsmtoiva,
		Bsmtofte:       dbDoc.Bsmtofte,
		RetmunMto:      dbDoc.RetmunMto,
		Dolarflete:     dbDoc.Dolarflete,
		Bsretflete:     dbDoc.Bsretflete,
		Dretflete:      dbDoc.Dretflete,
		DretmunMto:     dbDoc.DretmunMto,
		Retivaoblig:    dbDoc.Retivaoblig,
		Edoentrega:     dbDoc.Edoentrega,
		Dmtoiva:        dbDoc.Dmtoiva,
		Prcdctoaplic:   dbDoc.Prcdctoaplic,
		Montodctodol:   dbDoc.Montodctodol,
		Montodctobs:    dbDoc.Montodctobs,
		CreatedAt:      dbDoc.CreatedAt,
		UpdatedAt:      dbDoc.UpdatedAt,
		DeletedAt:      dbDoc.DeletedAt.Time,
	}
}

func DbDocToDocLine(dbDoc *db.AdminGetDocumentsWithLinesRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       dbDoc.Agencia_2.String,
		Tipodoc:       dbDoc.Tipodoc_2.String,
		Documento:     dbDoc.Documento_2.String,
		Tipodocv:      dbDoc.Tipodocv_2.String,
		Grupo:         dbDoc.Grupo.Int32,
		Subgrupo:      dbDoc.Subgrupo.Int32,
		Origen:        dbDoc.Origen.String,
		Codigo:        dbDoc.Codigo.String,
		Codhijo:       dbDoc.Codhijo.String,
		Pid:           dbDoc.Pid.String,
		Nombre:        dbDoc.Nombre.String,
		Cantidad:      dbDoc.Cantidad.Int32,
		Cntdevuelt:    dbDoc.Cntdevuelt.Int32,
		Vndcntdevuelt: dbDoc.Vndcntdevuelt.String,
		Dvndmtototal:  dbDoc.Dvndmtototal_2.String,
		Dpreciofin:    dbDoc.Dpreciofin.String,
		Dpreciounit:   dbDoc.Dpreciounit.String,
		Dmontoneto:    dbDoc.Dmontoneto.String,
		Dmontototal:   dbDoc.Dmontototal.String,
		Timpueprc:     dbDoc.Timpueprc.String,
		Unidevuelt:    dbDoc.Unidevuelt.Int32,
		Fechadoc:      dbDoc.Fechadoc.Time,
		Vendedor:      dbDoc.Vendedor_2.String,
		Codcoord:      dbDoc.Codcoord_2.String,
		CreatedAt:     dbDoc.CreatedAt_2.Time,
		UpdatedAt:     dbDoc.UpdatedAt_2.Time,
		DeletedAt:     dbDoc.DeletedAt_2.Time,
	}
}

func DbDocByCodeToDocument(dbDoc *db.GetDocumentsWithLinesBySalesmanRow) *Document {
	return &Document{
		Agencia:        dbDoc.Agencia,
		Tipodoc:        dbDoc.Tipodoc,
		Documento:      dbDoc.Documento,
		Tipodocv:       dbDoc.Tipodocv,
		Codcliente:     dbDoc.Codcliente,
		Nombrecli:      dbDoc.Nombrecli,
		Contribesp:     dbDoc.Contribesp,
		RutaParme:      dbDoc.RutaParme,
		Tipoprecio:     dbDoc.Tipoprecio,
		Emision:        dbDoc.Emision,
		Recepcion:      dbDoc.Recepcion,
		Vence:          dbDoc.Vence,
		Diascred:       dbDoc.Diascred,
		Estatusdoc:     dbDoc.Estatusdoc,
		Dtotneto:       dbDoc.Dtotneto,
		Dtotimpuest:    dbDoc.Dtotimpuest,
		Dtotalfinal:    dbDoc.Dtotalfinal,
		Dtotpagos:      dbDoc.Dtotpagos,
		Dtotdescuen:    dbDoc.Dtotdescuen,
		Dflete:         dbDoc.Dflete,
		Dtotdev:        dbDoc.Dtotdev,
		Dvndmtototal:   dbDoc.Dvndmtototal,
		Dretencion:     dbDoc.Dretencion,
		Dretencioniva:  dbDoc.Dretencioniva,
		Vendedor:       dbDoc.Vendedor,
		Codcoord:       dbDoc.Codcoord,
		Aceptadev:      dbDoc.Aceptadev,
		KtiNegesp:      dbDoc.KtiNegesp,
		Bsiva:          dbDoc.Bsiva,
		Bsflete:        dbDoc.Bsflete,
		Bsretencion:    dbDoc.Bsretencion,
		Bsretencioniva: dbDoc.Bsretencioniva,
		Tasadoc:        dbDoc.Tasadoc,
		Mtodcto:        dbDoc.Mtodcto,
		Fchvencedcto:   dbDoc.Fchvencedcto,
		Tienedcto:      dbDoc.Tienedcto,
		Cbsret:         dbDoc.Cbsret,
		Cdret:          dbDoc.Cdret,
		Cbsretiva:      dbDoc.Cbsretiva,
		Cdretiva:       dbDoc.Cdretiva,
		Cbsrparme:      dbDoc.Cbsrparme,
		Cdrparme:       dbDoc.Cdrparme,
		Cbsretflete:    dbDoc.Cbsretflete,
		Cdretflete:     dbDoc.Cdretflete,
		Bsmtoiva:       dbDoc.Bsmtoiva,
		Bsmtofte:       dbDoc.Bsmtofte,
		RetmunMto:      dbDoc.RetmunMto,
		Dolarflete:     dbDoc.Dolarflete,
		Bsretflete:     dbDoc.Bsretflete,
		Dretflete:      dbDoc.Dretflete,
		DretmunMto:     dbDoc.DretmunMto,
		Retivaoblig:    dbDoc.Retivaoblig,
		Edoentrega:     dbDoc.Edoentrega,
		Dmtoiva:        dbDoc.Dmtoiva,
		Prcdctoaplic:   dbDoc.Prcdctoaplic,
		Montodctodol:   dbDoc.Montodctodol,
		Montodctobs:    dbDoc.Montodctobs,
		CreatedAt:      dbDoc.CreatedAt,
		UpdatedAt:      dbDoc.UpdatedAt,
		DeletedAt:      dbDoc.DeletedAt.Time,
	}
}

func DbDocByCodeToDocLine(dbDoc *db.GetDocumentsWithLinesBySalesmanRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       dbDoc.Agencia_2.String,
		Tipodoc:       dbDoc.Tipodoc_2.String,
		Documento:     dbDoc.Documento_2.String,
		Tipodocv:      dbDoc.Tipodocv_2.String,
		Grupo:         dbDoc.Grupo.Int32,
		Subgrupo:      dbDoc.Subgrupo.Int32,
		Origen:        dbDoc.Origen.String,
		Codigo:        dbDoc.Codigo.String,
		Codhijo:       dbDoc.Codhijo.String,
		Pid:           dbDoc.Pid.String,
		Nombre:        dbDoc.Nombre.String,
		Cantidad:      dbDoc.Cantidad.Int32,
		Cntdevuelt:    dbDoc.Cntdevuelt.Int32,
		Vndcntdevuelt: dbDoc.Vndcntdevuelt.String,
		Dvndmtototal:  dbDoc.Dvndmtototal_2.String,
		Dpreciofin:    dbDoc.Dpreciofin.String,
		Dpreciounit:   dbDoc.Dpreciounit.String,
		Dmontoneto:    dbDoc.Dmontoneto.String,
		Dmontototal:   dbDoc.Dmontototal.String,
		Timpueprc:     dbDoc.Timpueprc.String,
		Unidevuelt:    dbDoc.Unidevuelt.Int32,
		Fechadoc:      dbDoc.Fechadoc.Time,
		Vendedor:      dbDoc.Vendedor_2.String,
		Codcoord:      dbDoc.Codcoord_2.String,
		CreatedAt:     dbDoc.CreatedAt,
		UpdatedAt:     dbDoc.UpdatedAt,
		DeletedAt:     dbDoc.DeletedAt.Time,
	}
}
