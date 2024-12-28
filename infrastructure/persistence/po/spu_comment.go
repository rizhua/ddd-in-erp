package po

import "time"

// 商品评价
type SpuComment struct {
	ID         int64     `xorm:"id pk autoincr"`
	SpuID      int32     `xorm:"spu_id bigint"`
	Content    string    `xorm:"content varchar(42)"`
	Media      string    `xorm:"media text"`
	UserID     int32     `xorm:"user_id int"`
	UserAvatar string    `xorm:"user_avatar text"`
	Nickname   string    `xorm:"nickname varchar(30)"`
	LoveCount  int32     `xorm:"love_count int"`
	StarCount  int8      `xorm:"star_count int"`
	ReadCount  int32     `xorm:"read_count int"`
	OrgID      int32     `xorm:"org_id int"`
	Status     int8      `xorm:"status int"`
	CreateAt   time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *SpuComment) TableName() string {
	return "spu_comment"
}
