package command

import "errors"

// 创建部门
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

// 删除部门
type DeleteDept struct {
	ID []int64 `json:"id" binding:"required"`
}

func (t *DeleteDept) Validate() error {
	if len(t.ID) == 0 {
		return errors.New("部门ID不能为空")
	}
	return nil
}

// 更新部门
type UpdateDept struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentID int64  `json:"parentId"`
	MgrID    int64  `json:"mgrId"`
	OrgID    int64
}

func (t *UpdateDept) Validate() error {
	if t.ID == 0 {
		return errors.New("部门ID不能为空")
	}
	if t.Name == "" {
		return errors.New("部门名称不能为空")
	}
	return nil
}
