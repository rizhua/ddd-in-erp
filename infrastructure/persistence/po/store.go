package po

import "time"

// 店仓
type Store struct {
	ID       int32         `xorm:"pk autoincr id" json:"id"`
	Name     string        `xorm:"varchar(128) name" json:"name"`
	Icon     string        `xorm:"text icon" json:"icon"`
	OrgID    int32         `xorm:"int org_id" json:"orgId"`
	Lng      float32       `xorm:"float lng" json:"lng"`
	Lat      float32       `xorm:"float lat" json:"lat"`
	Contact  string        `xorm:"varchar(64) contact" json:"contact"`
	Tel      string        `xorm:"tel varchar(16)" json:"tel"`
	Address  string        `xorm:"address varchar(512)" json:"address"`
	OpenTime string        `xorm:"open_time datetime" json:"openTime"`
	StopTime string        `xorm:"stop_time datetime" json:"stopTime"`
	Place    []interface{} `xorm:"json place" json:"place"`
	Type     int8          `xorm:"type int" json:"type"`
	Status   int8          `xorm:"int status" json:"status"`
	UpdateAt time.Time     `xorm:"update_at datetime" json:"updateAt"`
	CreateAt time.Time     `xorm:"create_at datetime" json:"createAt"`
}

// 设定表名
func (t *Store) TableName() string {
	return "store"
}
