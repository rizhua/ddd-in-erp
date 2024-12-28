package po

import "time"

// 退换/售后
type Refund struct {
	ID         int32         `xorm:"id pk autoincr"`
	RefundNo   string        `xorm:"refund_no varchar(32) notnull refund_no"`
	Data       []interface{} `xorm:"data json notnull"`
	Reason     string        `xorm:"reason varchar(16)"`
	Detail     string        `xorm:"detail varchar(420) notnull"`
	Amount     int32         `xorm:"amount int notnull"`
	VerifyTime time.Time     `xorm:"verify_time datetime notnull"`
	VerifyText string        `xorm:"verify_text varchar(420) notnull"`
	Status     int8          `xorm:"status int notnull"`
	OrderNo    string        `xorm:"order_no varchar(32) notnull"`
	UpdateAt   time.Time     `xorm:"update_at datetime notnull"`
	CreateAt   time.Time     `xorm:"create_at datetime notnull"`
}

// 设定表名
func (t *Refund) TableName() string {
	return "refund"
}
