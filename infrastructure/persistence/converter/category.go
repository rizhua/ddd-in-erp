package converter

import (
	"rizhua.com/infrastructure/persistence/po"
)

// 商品类目
type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	Path     string `json:"path"`
	Sort     int8   `json:"sort"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Category) Info(in po.Category) (info Category) {
	info = Category{
		ID:       in.ID,
		Name:     in.Name,
		ParentID: in.ParentID,
		Path:     in.Path,
		Sort:     in.Sort,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Category) List(in []po.Category) (list []Category) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}

type CategoryTree struct {
	Category
	Leaf     bool           `json:"leaf"`
	Children []CategoryTree `json:"children"`
}

func (t *Category) Tree(in []po.Category, parentID int64) (list []CategoryTree) {
	for _, v := range in {
		if parentID == v.ParentID {
			temp := CategoryTree{
				Category: t.Info(v),
				Leaf:     false,
			}

			child := t.Tree(in, v.ID)
			if len(child) == 0 {
				temp.Leaf = true
				temp.Children = make([]CategoryTree, 0)
			} else {
				temp.Leaf = false
				temp.Children = child
			}
			list = append(list, temp)
		}
	}
	return
}

type CategoryAttribute struct {
	ID         int64    `json:"id"`
	Label      string   `json:"label"`
	Value      []string `json:"value"`
	Type       string   `json:"type"`
	CategoryID int64    `json:"categoryId"`
	Required   bool     `json:"required"`
	UpdateAt   string   `json:"updateAt"`
	CreateAt   string   `json:"createAt"`
}

func (t *CategoryAttribute) Info(in po.CategoryAttribute) (info CategoryAttribute) {
	info = CategoryAttribute{
		ID:         in.ID,
		Label:      in.Label,
		Value:      in.Value,
		Type:       in.Type,
		CategoryID: in.CategoryID,
		Required:   in.Required,
		UpdateAt:   in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt:   in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *CategoryAttribute) List(in []po.CategoryAttribute) (list []CategoryAttribute) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
