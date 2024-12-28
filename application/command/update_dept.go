package command

type UpdateDept struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentId"`
	MgrID    int64  `json:"mgrId"`
	OrgID    int64
}
