package po

import "time"

// 购物车
type Cart struct {
	ID       int64     `xorm:"id pk autoincr"`
	SpuID    int64     `xorm:"spu_id varchar(30)"`
	SkuID    int64     `xorm:"sku_id int"`
	Number   int32     `xorm:"qty int"`
	UserID   int64     `xorm:"user_id bigint"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Cart) TableName() string {
	return "Cart"
}
