package po

type RoleNode struct {
	RoleID int64 `xorm:"role_id int"`
	NodeID int64 `xorm:"node_id int"`
}

// 设定表名
func (t *RoleNode) TableName() string {
	return "role_node"
}
