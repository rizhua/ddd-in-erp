package po

import (
	"time"

	"rizhua.com/pkg/util"
)

// 组织
type Org struct {
	ID       int64     `xorm:"id pk autoincr"`
	Code     string    `xorm:"code varchar(16)"`
	Name     string    `xorm:"name varchar(128)"`
	FullName string    `xorm:"full_name varchar(128)"`
	Icon     string    `xorm:"icon varchar(1024)"`
	Industry string    `xorm:"industry varchar(128)"`
	Capacity int8      `xorm:"capacity int"`
	Contact  string    `xorm:"contact varchar(128)"`
	Tel      string    `xorm:"tel varchar(16)"`
	Address  string    `xorm:"address varchar(1024)"`
	License  string    `xorm:"license varchar(256)"`
	OwnerID  int64     `xorm:"owner_id int"`
	Status   int8      `xorm:"status int"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Org) TableName() string {
	return "org"
}

func (t *Org) BeforeInsert() {
	t.ID = util.NewMist().Generate()
	t.Status = 1
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
	if t.Code == "" {
		t.Code = util.RandomString(8)
	}
}

func (t *Org) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
