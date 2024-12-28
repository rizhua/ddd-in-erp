package po

import "time"

// 代理商申请表
type AgentApply struct {
	ID     int64     `xorm:"id pk autoincr"`
	BizID  int64     `xorm:"biz_id bigint"`
	AText  string    `xorm:"a_text varchar(1024)"`
	ATime  time.Time `xorm:"a_time datetime"`
	VText  string    `xorm:"v_text varchar(1024)"`
	VTime  time.Time `xorm:"v_time datetime"`
	Type   int8      `xorm:"type int"`
	Status int8      `xorm:"status int"`
	OrgID  int64     `xorm:"org_id bigint"`
}

// 设定表名
func (t *AgentApply) TableName() string {
	return "agent_apply"
}
