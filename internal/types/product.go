package types

import (
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `json:"id"`
	Grupo        int32     `json:"grupo"`
	Subgrupo     int32     `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Codigo       string    `json:"codigo"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   int32     `json:"existencia"`
	Precio1      string    `json:"precio1"`
	Precio2      string    `json:"precio2"`
	Precio3      string    `json:"precio3"`
	Precio4      string    `json:"precio4"`
	Precio5      string    `json:"precio5"`
	Precio6      string    `json:"precio6"`
	Precio7      string    `json:"precio7"`
	Discont      bool      `json:"discont"`
	VtaMax       int32     `json:"vta_max"`
	VtaMin       int32     `json:"vta_min"`
	Dctotope     string    `json:"dctotope"`
	Enpreventa   bool      `json:"enpreventa"`
	Comprometido int32     `json:"comprometido"`
	VtaMinenx    int32     `json:"vta_minenx"`
	VtaSolofac   bool      `json:"vta_solofac"`
	VtaSolone    bool      `json:"vta_solone"`
	Codbarras    string    `json:"codbarras"`
	Detalles     string    `json:"detalles"`
	Cantbulto    int32     `json:"cantbulto"`
	CostoProm    string    `json:"costoProm"`
	Util1        string    `json:"util1"`
	Util2        string    `json:"util2"`
	Util3        string    `json:"util3"`
	Fchultcomp   time.Time `json:"fchultcomp"`
	Qtyultcomp   int32     `json:"qtyultcomp"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"-"`
}

type CreateProductRequest struct {
	Grupo        int32     `json:"grupo"`
	Subgrupo     int32     `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Codigo       string    `json:"codigo"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   int32     `json:"existencia"`
	Precio1      string    `json:"precio1"`
	Precio2      string    `json:"precio2"`
	Precio3      string    `json:"precio3"`
	Precio4      string    `json:"precio4"`
	Precio5      string    `json:"precio5"`
	Precio6      string    `json:"precio6"`
	Precio7      string    `json:"precio7"`
	Discont      bool      `json:"discont"`
	VtaMax       int32     `json:"vta_max"`
	VtaMin       int32     `json:"vta_min"`
	Dctotope     string    `json:"dctotope"`
	Enpreventa   bool      `json:"enpreventa"`
	Comprometido int32     `json:"comprometido"`
	VtaMinenx    int32     `json:"vta_minenx"`
	VtaSolofac   bool      `json:"vta_solofac"`
	VtaSolone    bool      `json:"vta_solone"`
	Codbarras    string    `json:"codbarras"`
	Detalles     string    `json:"detalles"`
	Cantbulto    int32     `json:"cantbulto"`
	CostoProm    string    `json:"costoProm"`
	Util1        string    `json:"util1"`
	Util2        string    `json:"util2"`
	Util3        string    `json:"util3"`
	Fchultcomp   time.Time `json:"fchultcomp"`
	Qtyultcomp   int32     `json:"qtyultcomp"`
}

type UpdateProductRequest struct {
	Grupo        int32     `json:"grupo"`
	Subgrupo     int32     `json:"subgrupo"`
	Nombre       string    `json:"nombre"`
	Referencia   string    `json:"referencia"`
	Marca        string    `json:"marca"`
	Unidad       string    `json:"unidad"`
	Existencia   int32     `json:"existencia"`
	Precio1      string    `json:"precio1"`
	Precio2      string    `json:"precio2"`
	Precio3      string    `json:"precio3"`
	Precio4      string    `json:"precio4"`
	Precio5      string    `json:"precio5"`
	Precio6      string    `json:"precio6"`
	Precio7      string    `json:"precio7"`
	Discont      bool      `json:"discont"`
	VtaMax       int32     `json:"vta_max"`
	VtaMin       int32     `json:"vta_min"`
	Dctotope     string    `json:"dctotope"`
	Enpreventa   bool      `json:"enpreventa"`
	Comprometido int32     `json:"comprometido"`
	VtaMinenx    int32     `json:"vta_minenx"`
	VtaSolofac   bool      `json:"vta_solofac"`
	VtaSolone    bool      `json:"vta_solone"`
	Codbarras    string    `json:"codbarras"`
	Detalles     string    `json:"detalles"`
	Cantbulto    int32     `json:"cantbulto"`
	CostoProm    string    `json:"costoProm"`
	Util1        string    `json:"util1"`
	Util2        string    `json:"util2"`
	Util3        string    `json:"util3"`
	Fchultcomp   time.Time `json:"fchultcomp"`
	Qtyultcomp   int32     `json:"qtyultcomp"`
	ID           uuid.UUID `json:"id"`
}

func DbProductToProduct(db *db.Articulo) (*Product, error) {
	id, err := uuid.Parse(db.ID)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:           id,
		Grupo:        db.Grupo,
		Subgrupo:     db.Subgrupo,
		Nombre:       db.Nombre,
		Codigo:       db.Codigo,
		Referencia:   db.Referencia,
		Marca:        db.Marca,
		Unidad:       db.Unidad,
		Existencia:   db.Existencia,
		Precio1:      db.Precio1,
		Precio2:      db.Precio2,
		Precio3:      db.Precio3,
		Precio4:      db.Precio4,
		Precio5:      db.Precio5,
		Precio6:      db.Precio6,
		Precio7:      db.Precio7,
		Discont:      db.Discont,
		VtaMax:       db.VtaMax,
		VtaMin:       db.VtaMin,
		Dctotope:     db.Dctotope,
		Enpreventa:   db.Enpreventa,
		Comprometido: db.Comprometido,
		VtaMinenx:    db.VtaMinenx,
		VtaSolofac:   db.VtaSolofac,
		VtaSolone:    db.VtaSolone,
		Codbarras:    db.Codbarras,
		Detalles:     db.Detalles,
		Cantbulto:    db.Cantbulto,
		CostoProm:    db.CostoProm,
		Util1:        db.Util1,
		Util2:        db.Util2,
		Util3:        db.Util3,
		Fchultcomp:   db.Fchultcomp,
		Qtyultcomp:   db.Qtyultcomp,
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		DeletedAt:    db.DeletedAt.Time,
	}, nil
}

func CreateProductToDb(r *CreateProductRequest) (*db.CreateProductParams, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &db.CreateProductParams{
		ID:           id.String(),
		Grupo:        r.Grupo,
		Subgrupo:     r.Subgrupo,
		Nombre:       r.Nombre,
		Codigo:       r.Codigo,
		Referencia:   r.Referencia,
		Marca:        r.Marca,
		Unidad:       r.Unidad,
		Existencia:   r.Existencia,
		Precio1:      r.Precio1,
		Precio2:      r.Precio2,
		Precio3:      r.Precio3,
		Precio4:      r.Precio4,
		Precio5:      r.Precio5,
		Precio6:      r.Precio6,
		Precio7:      r.Precio7,
		Discont:      r.Discont,
		VtaMax:       r.VtaMax,
		VtaMin:       r.VtaMin,
		Dctotope:     r.Dctotope,
		Enpreventa:   r.Enpreventa,
		Comprometido: r.Comprometido,
		VtaMinenx:    r.VtaMinenx,
		VtaSolofac:   r.VtaSolofac,
		VtaSolone:    r.VtaSolone,
		Codbarras:    r.Codbarras,
		Detalles:     r.Detalles,
		Cantbulto:    r.Cantbulto,
		CostoProm:    r.CostoProm,
		Util1:        r.Util1,
		Util2:        r.Util2,
		Util3:        r.Util3,
		Fchultcomp:   r.Fchultcomp,
		Qtyultcomp:   r.Qtyultcomp,
	}, nil
}

func UpdateProductToDb(r *UpdateProductRequest) *db.UpdateProductParams {
	return &db.UpdateProductParams{
		Grupo:        r.Grupo,
		Subgrupo:     r.Subgrupo,
		Nombre:       r.Nombre,
		Referencia:   r.Referencia,
		Marca:        r.Marca,
		Unidad:       r.Unidad,
		Existencia:   r.Existencia,
		Precio1:      r.Precio1,
		Precio2:      r.Precio2,
		Precio3:      r.Precio3,
		Precio4:      r.Precio4,
		Precio5:      r.Precio5,
		Precio6:      r.Precio6,
		Precio7:      r.Precio7,
		Discont:      r.Discont,
		VtaMax:       r.VtaMax,
		VtaMin:       r.VtaMin,
		Dctotope:     r.Dctotope,
		Enpreventa:   r.Enpreventa,
		Comprometido: r.Comprometido,
		VtaMinenx:    r.VtaMinenx,
		VtaSolofac:   r.VtaSolofac,
		VtaSolone:    r.VtaSolone,
		Codbarras:    r.Codbarras,
		Detalles:     r.Detalles,
		Cantbulto:    r.Cantbulto,
		CostoProm:    r.CostoProm,
		Util1:        r.Util1,
		Util2:        r.Util2,
		Util3:        r.Util3,
		Fchultcomp:   r.Fchultcomp,
		Qtyultcomp:   r.Qtyultcomp,
		ID:           r.ID.String(),
	}
}
