package po

// 库存量单位
type Sku struct {
	ID       int64  `xorm:"id pk autoincr"`
	SpuID    int64  `xorm:"spu_id int"`
	Price    int16  `xorm:"price int"`
	Stock    int16  `xorm:"stock int"`
	Barcode  string `xorm:"barcode varchar(32)"`
	Discount int    `xorm:"discount int"`
	ShopID   int32  `xorm:"shop_id"`
}

// 设定表名
func (t *Sku) TableName() string {
	return "sku"
}
