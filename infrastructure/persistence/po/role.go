package po

import "time"

// 角色
type Role struct {
	ID       int64     `xorm:"pk autoincr id"`
	Name     string    `xorm:"varchar(90) notnull name"`
	OrgID    int64     `xorm:"int org_id"`
	ParentID int64     `xorm:"int parent_id"`
	UpdateAt time.Time `xorm:"datatime update_at"`
	CreateAt time.Time `xorm:"datetime create_at"`
}

// 设定表名
func (t *Role) TableName() string {
	return "role"
}

func (t *Role) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Role) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
