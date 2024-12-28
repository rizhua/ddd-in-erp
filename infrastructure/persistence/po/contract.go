package po

import "time"

// 合同
type Contract struct {
	ID         int64     `xorm:"id pk autoincr"`
	Number     string    `xorm:"number varchar(32)"`
	CustomerID int32     `xorm:"customer_id int"`
	Content    string    `xorm:"content text"`
	Amount     int32     `xorm:"amount int"`
	PaidAmount int32     `xorm:"paid_amount int"`
	Begtime    time.Time `xorm:"begtime datetime"`
	Endtime    time.Time `xorm:"endtime datetime"`
	WeAgent    string    `xorm:"we_agent varchar(32)"`
	HeAgent    string    `xorm:"he_agent varchar(32)"`
	DueTime    time.Time `xorm:"duetime datetime"`
	Remark     string    `xorm:"remark varchar(420)"`
	Status     int8      `xorm:"status int"`
	OrgID      int32     `xorm:"org_id int"`
	UpdateAt   time.Time `xorm:"update_at datetime"`
	CreateAt   time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Contract) TableName() string {
	return "contract"
}
