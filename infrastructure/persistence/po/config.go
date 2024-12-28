package po

import "time"

// 配置
type Config struct {
	ID       int32                  `xorm:"id pk autoincr"`
	Code     string                 `xorm:"code varchar(128) notnull"`
	Data     map[string]interface{} `xorm:"data json"`
	Remark   string                 `xorm:"remark text"`
	UpdateAt time.Time              `xorm:"update_at datetime notnull"`
	CreateAt time.Time              `xorm:"create_at datetime notnull"`
}

// 设定表名
func (t *Config) TableName() string {
	return "config"
}

func (t *Config) BeforeInsert() {
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Config) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
