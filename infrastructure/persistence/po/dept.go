package po

import "time"

// 部门
type Dept struct {
	ID       int64     `xorm:"id pk autoincr"`
	Name     string    `xorm:"name varchar(90)"`
	OrgID    int64     `xorm:"org_id bigint"`
	ParentID int64     `xorm:"parent_id bigint"`
	MgrID    int64     `xorm:"mgr_id bigint"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Dept) TableName() string {
	return "dept"
}

func (t *Dept) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Dept) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
