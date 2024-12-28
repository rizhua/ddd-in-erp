package po

import "time"

// 采购、销售商品
type OrderItem struct {
	ID      int32     `xorm:"id pk autoincr"`
	OrderNO string    `xorm:"order_no varchar(32)"`
	SpuName string    `xorm:"spu_name varchar(30) notnull"`
	SpuMeta string    `xorm:"spu_meta json notnull"`
	Sku     string    `xorm:"sku text notnull"`
	Price   time.Time `xorm:"price int notnull"`
	Amount  int8      `xorm:"amount int notnull"`
	Deduct  string    `xorm:"deduct json notnull"`
}

// 设定表名
func (t *OrderItem) TableName() string {
	return "OrderItem"
}
