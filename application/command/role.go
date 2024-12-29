package command

// 创建角色
type CreateRole struct {
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentId"`
	OrgID    int64
}

// 更新角色
type UpdateRole struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	OrgID    int64
}

// 角色绑定、解绑用户
type BindingUser struct {
	RoleID int64   `json:"roleId" binding:"required"`
	UserID []int64 `json:"userId"`
}

// 角色绑定、解绑节点
type BindingNode struct {
	RoleID int64   `json:"roleId" binding:"required"`
	NodeID []int64 `json:"nodeId"`
}
