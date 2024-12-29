package command

import "errors"

// 创建类目
type CreateCategory struct {
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}

// 更新类目
type UpdateCategory struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}

// 创建类目属性
type CreateCategoryAttribute struct {
	CategoryID int64    `json:"categoryId"`
	Label      string   `json:"label"`
	Value      []string `json:"value"`
	Type       string   `json:"type"`
}

func (t *CreateCategoryAttribute) Validate() error {
	if t.Label == "" {
		return errors.New("参数错误，缺少参数 label")
	}
	return nil
}

// 删除类目属性
type DeleteCategoryAttribute struct {
	CategoryID int64   `json:"categoryId"`
	ID         []int64 `json:"id"`
}

func (c *DeleteCategoryAttribute) Validate() error {
	if len(c.ID) == 0 {
		return errors.New("参数错误，缺少参数 id")
	}
	return nil
}

// 更新类目属性
type UpdateCategoryAttribute struct {
	ID    int64    `json:"id"`
	Label string   `json:"label"`
	Value []string `json:"value"`
	Type  string   `json:"type"`
}

func (t *UpdateCategoryAttribute) Validate() error {
	if t.ID == 0 {
		return errors.New("参数错误，缺少参数 id")
	}
	return nil
}
