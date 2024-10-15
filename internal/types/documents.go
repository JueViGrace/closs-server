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
	Contribesp     string    `json:"contribesp"`
	RutaParme      string    `json:"ruta_parme"`
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
	Fechamodifi    time.Time `json:"fechamodifi"`
	Aceptadev      string    `json:"aceptadev"`
	KtiNegesp      string    `json:"kti_negesp"`
	Bsiva          string    `json:"bsiva"`
	Bsflete        string    `json:"bsflete"`
	Bsretencion    string    `json:"bsretencion"`
	Bsretencioniva string    `json:"bsretencioniva"`
	Tasadoc        string    `json:"tasadoc"`
	Mtodcto        string    `json:"mtodcto"`
	Fchvencedcto   time.Time `json:"fchvencedcto"`
	Tienedcto      string    `json:"tienedcto"`
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
	Dolarflete     int32     `json:"dolarflete"`
	Bsretflete     string    `json:"bsretflete"`
	Dretflete      string    `json:"dretflete"`
	DretmunMto     string    `json:"dretmun_mto"`
	Retivaoblig    uint8     `json:"retivaoblig"`
	Edoentrega     uint8     `json:"edoentrega"`
	Dmtoiva        string    `json:"dmtoiva"`
	Prcdctoaplic   string    `json:"prcdctoaplic"`
	Montodctodol   string    `json:"montodctodol"`
	Montodctobs    string    `json:"montodctobs"`
}

type DocumentLine struct {
	Agencia       string    `json:"agencia"`
	Tipodoc       string    `json:"tipodoc"`
	Documento     string    `json:"documento"`
	Tipodocv      string    `json:"tipodocv"`
	Grupo         string    `json:"grupo"`
	Subgrupo      string    `json:"subgrupo"`
	Origen        string    `json:"origen"`
	Codigo        string    `json:"codigo"`
	Codhijo       string    `json:"codhijo"`
	Pid           string    `json:"pid"`
	Nombre        string    `json:"nombre"`
	Cantidad      string    `json:"cantidad"`
	Cntdevuelt    string    `json:"cntdevuelt"`
	Vndcntdevuelt string    `json:"vndcntdevuelt"`
	Dvndmtototal  string    `json:"dvndmtototal"`
	Dpreciofin    string    `json:"dpreciofin"`
	Dpreciounit   string    `json:"dpreciounit"`
	Dmontoneto    string    `json:"dmontoneto"`
	Dmontototal   string    `json:"dmontototal"`
	Timpueprc     string    `json:"timpueprc"`
	Unidevuelt    string    `json:"unidevuelt"`
	Fechadoc      time.Time `json:"fechadoc"`
	Vendedor      string    `json:"vendedor"`
	Codcoord      string    `json:"codcoord"`
	Fechamodifi   time.Time `json:"fechamodifi"`
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
		Fechamodifi:    kd.Fechamodifi,
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
	}
}

func DbKeDoclToDocumentLine(kd *db.KeDoclmv) *DocumentLine {
	return &DocumentLine{
		Agencia:       kd.Agencia,
		Tipodoc:       kd.Tipodoc,
		Documento:     kd.Documento,
		Tipodocv:      kd.Tipodocv,
		Grupo:         kd.Grupo,
		Subgrupo:      kd.Subgrupo,
		Origen:        kd.Origen,
		Codigo:        kd.Codigo,
		Codhijo:       kd.Codhijo,
		Pid:           kd.Pid,
		Nombre:        kd.Nombre,
		Cantidad:      kd.Cantidad,
		Cntdevuelt:    kd.Cntdevuelt,
		Vndcntdevuelt: kd.Vndcntdevuelt,
		Dvndmtototal:  kd.Dvndmtototal,
		Dpreciofin:    kd.Dpreciofin,
		Dpreciounit:   kd.Dpreciounit,
		Dmontoneto:    kd.Dmontoneto,
		Dmontototal:   kd.Dmontototal,
		Timpueprc:     kd.Timpueprc,
		Unidevuelt:    kd.Unidevuelt,
		Fechadoc:      kd.Fechadoc,
		Vendedor:      kd.Vendedor,
		Codcoord:      kd.Codcoord,
		Fechamodifi:   kd.Fechamodifi,
	}
}

// func DbDocToDocumentWithLines(dbDocs []db.FindAllDocumentsWithLinesRow) []*DocumentWithLines {
// 	docs := make([]*DocumentWithLines, 0)
// 	docMap := make(map[Document][]DocumentLine)
// 	doc := new(Document)
//
// 	for _, dbDoc := range dbDocs {
// 		if doc == nil {
// 			doc = dbDocToDocument(dbDoc)
// 		}
//
// 		if doc.Documento != dbDoc.Documento {
// 			doc = dbDocToDocument(dbDoc)
// 		}
//
// 		docMap[*doc] = append(docMap[*doc], *dbDocToDocLine(dbDoc))
// 	}
//
// 	for key, value := range docMap {
// 		docs = append(docs, docMapToDocWithLines(&key, &value))
// 	}
//
// 	return docs
// }

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
			Fechamodifi:    key.Fechamodifi,
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
		},
		Lines: *value,
	}
}

func DbDocToDocument(dbDoc *db.FindAllDocumentsWithLinesRow) *Document {
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
		Fechamodifi:    dbDoc.Fechamodifi,
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
	}
}

func DbDocToDocLine(dbDoc *db.FindAllDocumentsWithLinesRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       dbDoc.Agencia_2.String,
		Tipodoc:       dbDoc.Tipodoc_2.String,
		Documento:     dbDoc.Documento_2.String,
		Tipodocv:      dbDoc.Tipodocv_2.String,
		Grupo:         dbDoc.Grupo.String,
		Subgrupo:      dbDoc.Subgrupo.String,
		Origen:        dbDoc.Origen.String,
		Codigo:        dbDoc.Codigo.String,
		Codhijo:       dbDoc.Codhijo.String,
		Pid:           dbDoc.Pid.String,
		Nombre:        dbDoc.Nombre.String,
		Cantidad:      dbDoc.Cantidad.String,
		Cntdevuelt:    dbDoc.Cntdevuelt.String,
		Vndcntdevuelt: dbDoc.Vndcntdevuelt.String,
		Dvndmtototal:  dbDoc.Dvndmtototal_2.String,
		Dpreciofin:    dbDoc.Dpreciofin.String,
		Dpreciounit:   dbDoc.Dpreciounit.String,
		Dmontoneto:    dbDoc.Dmontoneto.String,
		Dmontototal:   dbDoc.Dmontototal.String,
		Timpueprc:     dbDoc.Timpueprc.String,
		Unidevuelt:    dbDoc.Unidevuelt.String,
		Fechadoc:      dbDoc.Fechadoc.Time,
		Vendedor:      dbDoc.Vendedor_2.String,
		Codcoord:      dbDoc.Codcoord_2.String,
		Fechamodifi:   dbDoc.Fechamodifi_2.Time,
	}
}

func DbDocByCodeToDocument(dbDoc *db.FindAllDocumentsWithLinesByCodeRow) *Document {
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
		Fechamodifi:    dbDoc.Fechamodifi,
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
	}
}

func DbDocByCodeToDocLine(dbDoc *db.FindAllDocumentsWithLinesByCodeRow) *DocumentLine {
	return &DocumentLine{
		Agencia:       dbDoc.Agencia_2.String,
		Tipodoc:       dbDoc.Tipodoc_2.String,
		Documento:     dbDoc.Documento_2.String,
		Tipodocv:      dbDoc.Tipodocv_2.String,
		Grupo:         dbDoc.Grupo.String,
		Subgrupo:      dbDoc.Subgrupo.String,
		Origen:        dbDoc.Origen.String,
		Codigo:        dbDoc.Codigo.String,
		Codhijo:       dbDoc.Codhijo.String,
		Pid:           dbDoc.Pid.String,
		Nombre:        dbDoc.Nombre.String,
		Cantidad:      dbDoc.Cantidad.String,
		Cntdevuelt:    dbDoc.Cntdevuelt.String,
		Vndcntdevuelt: dbDoc.Vndcntdevuelt.String,
		Dvndmtototal:  dbDoc.Dvndmtototal_2.String,
		Dpreciofin:    dbDoc.Dpreciofin.String,
		Dpreciounit:   dbDoc.Dpreciounit.String,
		Dmontoneto:    dbDoc.Dmontoneto.String,
		Dmontototal:   dbDoc.Dmontototal.String,
		Timpueprc:     dbDoc.Timpueprc.String,
		Unidevuelt:    dbDoc.Unidevuelt.String,
		Fechadoc:      dbDoc.Fechadoc.Time,
		Vendedor:      dbDoc.Vendedor_2.String,
		Codcoord:      dbDoc.Codcoord_2.String,
		Fechamodifi:   dbDoc.Fechamodifi_2.Time,
	}
}
