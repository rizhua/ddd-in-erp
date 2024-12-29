package command

import "errors"

// 创建套餐
type CreateBundle struct {
	Name     string `json:"name" binding:"required"`
	Term     int    `json:"term"`
	Quota    int32  `json:"quota"`
	Price    int32  `json:"price"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *CreateBundle) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

// 删除套餐
type DeleteBundle struct {
	ID []int64 `json:"id" binding:"required"`
}

func (t *DeleteBundle) Validate() error {
	if len(t.ID) == 0 {
		return errors.New("id is required")
	}
	return nil
}

// 更新套餐
type UpdateBundle struct {
	ID    int64   `json:"id" binding:"required"`
	Name  string  `json:"name"`
	Term  int     `json:"term"`
	Quota int32   `json:"quota"`
	Price int32   `json:"price"`
	Node  []int64 `json:"node"`
}

func (t *UpdateBundle) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

// 套餐绑定、解绑节点
type BindBundle struct {
	BundleID int64   `json:"bundleId" binding:"required"`
	NodeID   []int64 `json:"nodeId"`
}
