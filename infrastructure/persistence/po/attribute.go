package po

import (
	"time"
)

// 商品属性
type Attribute struct {
	ID       int64     `xorm:"id pk autoincr"`
	Label    string    `xorm:"label varchar(64)"`
	Value    []string  `xorm:"value json"`
	Multi    bool      `xorm:"multi boolean"`
	Required bool      `xorm:"required boolean"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Attribute) TableName() string {
	return "attribute"
}

func (t *Attribute) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Attribute) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
