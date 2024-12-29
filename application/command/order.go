package command

import "errors"

type OrderType int8

const (
	// 套餐订单
	OrderTypePack OrderType = iota
	// 商品订单
	OrderTypeProduct
)

type UnifiedOrder struct {
	BizID int64     `json:"bizId"`
	Qty   int32     `json:"qty"`
	Type  OrderType `json:"type"`
}

func (t *UnifiedOrder) Validate() error {
	if t.BizID <= 0 {
		return errors.New("bizId必须大于0")
	}
	if t.Qty <= 0 {
		return errors.New("数量必须大于0")
	}
	return nil
}
