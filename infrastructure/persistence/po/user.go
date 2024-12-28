package po

import (
	"time"

	"github.com/google/uuid"
	"rizhua.com/pkg/util"
)

// 用户
type User struct {
	ID        int64     `xorm:"id pk autoincr"`
	Nickname  string    `xorm:"nickname varchar(32) notnull unique"`
	Avatar    string    `xorm:"avatar varchar(1024) notnull"`
	Mobile    string    `xorm:"mobile varchar(16) notnull"`
	Email     string    `xorm:"email varchar(128) notnull"`
	Password  string    `xorm:"password varchar(128) notnull"`
	Birthday  time.Time `xorm:"birthday date"`
	Gender    int8      `xorm:"gender int"`
	AccessKey string    `xorm:"access_key varchar(256)"`
	SecretKey string    `xorm:"secret_key varchar(256)"`
	Status    int8      `xorm:"status int default 1"`
	LastTime  time.Time `xorm:"last_time timestamp"`
	UpdateAt  time.Time `xorm:"update_at timestamp"`
	CreateAt  time.Time `xorm:"create_at timestamp"`
}

// 设定表名
func (t *User) TableName() string {
	return "user"
}

func (t *User) BeforeInsert() {
	dt := time.Now()
	t.ID = util.NewMist().Generate()
	if t.Nickname == "" {
		t.Nickname = util.RandomString(16)
	}
	t.AccessKey = uuid.NewString()
	t.SecretKey = uuid.NewString()
	t.Status = 1
	t.LastTime = dt
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *User) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
