package command

type CreateRole struct {
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentId"`
	OrgID    int64
}
