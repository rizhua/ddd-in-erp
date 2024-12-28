package po

import "time"

// 公文/公告
type Notice struct {
	ID       int64     `xorm:"id pk autoincr"`
	Title    string    `xorm:"title varchar(256) notnull"`
	Content  string    `xorm:"content text notnull"`
	Attach   string    `xorm:"attach text"`
	Scope    int8      `xorm:"scope int notnull"`
	Drafter  string    `xorm:"drafter varchar(64) notnull"`
	Type     int8      `xorm:"type int notnull"`
	OrgID    int64     `xorm:"org_id varchar(90) notnull"`
	UpdateAt time.Time `xorm:"update_at datetime notnull"`
	CreateAt time.Time `xorm:"create_at datetime notnull"`
}

// 设定表名
func (t *Notice) TableName() string {
	return "notice"
}

func (t *Notice) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Notice) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
