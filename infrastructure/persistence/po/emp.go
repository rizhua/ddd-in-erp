package po

import (
	"time"

	"rizhua.com/pkg/util"
)

// 员工
type Emp struct {
	ID       int64     `xorm:"id pk autoincr"`
	UserID   int64     `xorm:"user_id bigint"`
	Name     string    `xorm:"name varchar(128)"`
	Number   string    `xorm:"number varchar(8) notnull"`
	Gender   int8      `xorm:"gender int notnull"`
	Position string    `xorm:"position varchar(64)"`
	Grade    string    `xorm:"grade varchar(16)"`
	Tel      string    `xorm:"tel varchar(16)"`
	Email    string    `xorm:"email varchar(128)"`
	Address  string    `xorm:"address varchar(256)"`
	OrgID    int64     `xorm:"org_id bigint"`
	JoinTime time.Time `xorm:"join_time date"`
	QuitTime time.Time `xorm:"quit_time date"`
	Status   int8      `xorm:"status int"`
}

// 设定表名
func (t *Emp) TableName() string {
	return "emp"
}

func (t *Emp) BeforeInsert() {
	t.ID = util.NewMist().Generate()
}
