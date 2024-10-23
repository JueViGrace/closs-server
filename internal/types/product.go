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

func DbProductToProduct(p *db.Articulo) (*Product, error) {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:           id,
		Grupo:        p.Grupo,
		Subgrupo:     p.Subgrupo,
		Nombre:       p.Nombre,
		Codigo:       p.Codigo,
		Referencia:   p.Referencia,
		Marca:        p.Marca,
		Unidad:       p.Unidad,
		Existencia:   p.Existencia,
		Precio1:      p.Precio1,
		Precio2:      p.Precio2,
		Precio3:      p.Precio3,
		Precio4:      p.Precio4,
		Precio5:      p.Precio5,
		Precio6:      p.Precio6,
		Precio7:      p.Precio7,
		Discont:      p.Discont,
		VtaMax:       p.VtaMax,
		VtaMin:       p.VtaMin,
		Dctotope:     p.Dctotope,
		Enpreventa:   p.Enpreventa,
		Comprometido: p.Comprometido,
		VtaMinenx:    p.VtaMinenx,
		VtaSolofac:   p.VtaSolofac,
		VtaSolone:    p.VtaSolone,
		Codbarras:    p.Codbarras,
		Detalles:     p.Detalles,
		Cantbulto:    p.Cantbulto,
		CostoProm:    p.CostoProm,
		Util1:        p.Util1,
		Util2:        p.Util2,
		Util3:        p.Util3,
		Fchultcomp:   p.Fchultcomp,
		Qtyultcomp:   p.Qtyultcomp,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		DeletedAt:    p.DeletedAt.Time,
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
