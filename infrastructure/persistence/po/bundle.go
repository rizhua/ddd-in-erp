package po

import "time"

// 资源套餐
type Bundle struct {
	ID       int64     `xorm:"pk autoincr id"`
	Name     string    `xorm:"name varchar(90)"`
	Term     int       `xorm:"term int"`
	Quota    int32     `xorm:"quota int"`
	Price    int32     `xorm:"price int"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Bundle) TableName() string {
	return "bundle"
}

func (t *Bundle) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Bundle) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
