package po

import (
	"time"
)

// 许可
type License struct {
	ID       int64     `xorm:"pk autoincr id"`
	Code     string    `xorm:"code varchar(128)"`
	BizID    int64     `xorm:"biz_id int"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *License) TableName() string {
	return "license"
}

func (t *License) BeforeInsert() {
	now := time.Now()
	t.CreateAt = now
}
