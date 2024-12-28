package assembler

import "rizhua.com/domain/entity"

type Spu struct {
	ID        int64            `json:"id"`
	Code      string           `json:"code"`
	Name      string           `json:"name"`
	LowPrice  int32            `json:"lowPrice"`
	SaleCount int32            `json:"saleCount"`
	RateCount int32            `json:"rateCount"`
	Barcode   string           `json:"barcode"`
	Media     []map[string]any `json:"media"`
	Detail    string           `json:"detail"`
	UpdateAt  string           `json:"updateAt"`
	CreateAt  string           `json:"createAt"`
}

func (t *Spu) Info(in entity.Spu) (info Spu) {
	info = Spu{
		ID:        in.ID,
		Code:      in.Code,
		Name:      in.Name,
		LowPrice:  in.LowPrice,
		SaleCount: in.SaleCount,
		RateCount: in.RateCount,
		Barcode:   in.Barcode,
		Media:     in.Media,
		Detail:    in.Detail,
		// CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
		// UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Spu) List(in []entity.Spu) (list []Spu) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
