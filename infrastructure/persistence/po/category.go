package po

import "time"

// 发布分类
type Category struct {
	ID       int64     `xorm:"id pk autoincr"`
	Name     string    `xorm:"name varchar(128) notnull"`
	ParentID int64     `xorm:"parent_id bigint notnull"`
	Path     string    `xorm:"path text notnull"`
	Sort     int8      `xorm:"sort int notnull"`
	UpdateAt time.Time `xorm:"datetime update_at"`
	CreateAt time.Time `xorm:"datetime create_at"`
}

// 设定表名
func (t *Category) TableName() string {
	return "category"
}

func (t *Category) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Category) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
