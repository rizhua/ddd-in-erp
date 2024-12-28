package po

type RoleUser struct {
	RoleID int64 `xorm:"role_id int"`
	UserID int64 `xorm:"user_id int"`
}

// 设定表名
func (t *RoleUser) TableName() string {
	return "role_user"
}
