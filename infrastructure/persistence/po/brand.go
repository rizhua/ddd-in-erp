package po

import "time"

// 品牌
type Brand struct {
	ID       int64     `xorm:"id pk autoincr"`
	Name     string    `xorm:"name varchar(90) notnull"`
	Logo     string    `xorm:"logo text"`
	OrgID    int64     `xorm:"org_id bigint"`
	Deleted  time.Time `xorm:"deleted datetime"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Brand) TableName() string {
	return "brand"
}

func (t *Brand) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Brand) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
