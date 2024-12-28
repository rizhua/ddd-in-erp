package po

import (
	"time"
)

// 节点
type Node struct {
	ID       int64     `xorm:"pk autoincr id"`
	Icon     string    `xorm:"icon text"`
	Name     string    `xorm:"name varchar(128)"`
	Meta     string    `xorm:"meta text"`
	Type     int8      `xorm:"type int"`
	ParentID int64     `xorm:"parent_id int"`
	Path     string    `xorm:"path int"`
	Sort     int8      `xorm:"sort int"`
	Status   int8      `xorm:"status int"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Node) TableName() string {
	return "node"
}

func (t *Node) BeforeInsert() {
	t.Status = 1
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Node) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
