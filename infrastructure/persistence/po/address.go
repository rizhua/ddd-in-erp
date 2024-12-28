package po

import (
	"time"
)

// 地址: 收、发货地址
type Address struct {
	ID       int64     `xorm:"id pk autoincr" json:"id"`
	Tag      string    `xorm:"tag varchar(64)" json:"tag"`
	Contact  string    `xorm:"contact varchar(64)" json:"contact" validate:"required"`
	Region   string    `xorm:"region varchar(128)" json:"region" validate:"required"`
	Detail   string    `xorm:"detail varchar(128)" json:"detail" validate:"required"`
	Tel      string    `xorm:"tel varchar(16)" json:"tel"`
	Default  bool      `xorm:"default bool" json:"default"`
	UserID   int64     `xorm:"user_id int" json:"userId"`
	OrgID    int64     `xorm:"org_id int" json:"orgId"`
	UpdateAt time.Time `xorm:"update_at datetime" json:"updateAt"`
	CreateAt time.Time `xorm:"create_at datetime" json:"createAt"`
}

// 设定表名
func (t *Address) TableName() string {
	return "address"
}

func (t *Address) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Address) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
