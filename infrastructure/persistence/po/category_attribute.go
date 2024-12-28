package po

import "time"

// 类目属性
type CategoryAttribute struct {
	ID         int64     `xorm:"id pk autoincr"`
	Label      string    `xorm:"label varchar(64)"`
	Value      []string  `xorm:"value json"`
	Type       string    `xorm:"type varchar(16)"`
	CategoryID int64     `xorm:"category_id bigint"`
	UpdateAt   time.Time `xorm:"update_at datetime"`
	CreateAt   time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *CategoryAttribute) TableName() string {
	return "category_attribute"
}

func (t *CategoryAttribute) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *CategoryAttribute) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
