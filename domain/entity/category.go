package entity

import (
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var CategoryRepo repository.Category

// 根实体:商品类目
type Category struct {
	ID        int64
	Name      string
	ParentID  int64
	Path      string
	Sort      int8
	Attribute *categoryAttribute
}

func (c *Category) Create() error {
	catalog := po.Category{
		Name:     c.Name,
		ParentID: c.ParentID,
		Sort:     c.Sort,
	}
	return CategoryRepo.Create(catalog)
}

func (c *Category) Update() error {
	catalog := po.Category{
		ID:       c.ID,
		Name:     c.Name,
		ParentID: c.ParentID,
		Sort:     c.Sort,
	}
	return CategoryRepo.Update(catalog)
}

// 实体:类目属性
type categoryAttribute struct {
	ID         int64
	CategoryID int64
	Label      string
	Value      []string
	Type       string
	Required   bool
}
