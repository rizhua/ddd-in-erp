package command

import "errors"

type CreateDept struct {
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentId"`
	MgrID    int64  `json:"mgrId"`
}

func (t *CreateDept) Validate() error {
	if t.Name == "" {
		return errors.New("部门名称不能为空")
	}
	return nil
}
