package command

import "errors"

type UpdateNode struct {
	ID       int64  `json:"id" binding:"required"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Meta     string `json:"meta"`
	Type     int8   `json:"type"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}

func (t *UpdateNode) Validate() error {
	if t.ID == 0 {
		return errors.New("参数错误: 缺少参数 id")
	}

	if t.Name == "" {
		return errors.New("参数错误: 缺少参数 name")
	}

	if t.Meta == "" {
		return errors.New("参数错误: 缺少参数 meta")
	}

	return nil
}

type UpdateNodeSort struct {
	ID   int64 `json:"id" binding:"required"`
	Sort int8  `json:"sort"`
}

type UpdateNodeStatus struct {
	ID     int64 `json:"id" binding:"required"`
	Status int8  `json:"status"`
}
