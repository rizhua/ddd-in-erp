package command

type UpdateRole struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	OrgID    int64
}

type BindingUser struct {
	RoleID int64   `json:"roleId" binding:"required"`
	UserID []int64 `json:"userId"`
}

type BindingNode struct {
	RoleID int64   `json:"roleId" binding:"required"`
	NodeID []int64 `json:"nodeId"`
}
