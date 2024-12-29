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
	Attribute *CategoryAttribute
}

func (t *Category) Create() error {
	category := po.Category{
		Name:     t.Name,
		ParentID: t.ParentID,
		Sort:     t.Sort,
	}
	return CategoryRepo.Create(category)
}

func (t *Category) Update() error {
	category := po.Category{
		ID:       t.ID,
		Name:     t.Name,
		ParentID: t.ParentID,
		Sort:     t.Sort,
	}
	return CategoryRepo.Update(category)
}

// 实体:类目属性
type CategoryAttribute struct {
	ID    int64
	Label string
	Value []string
	Type  string
}

func (t *Category) CreateAttribute() error {
	categoryAttribute := po.CategoryAttribute{
		CategoryID: t.ID,
		Label:      t.Attribute.Label,
		Value:      t.Attribute.Value,
		Type:       t.Attribute.Type,
	}
	return CategoryRepo.CreateAttribute(categoryAttribute)
}

func (t *Category) UpdateAttribute() error {
	categoryAttribute := po.CategoryAttribute{
		ID:    t.Attribute.ID,
		Label: t.Attribute.Label,
		Value: t.Attribute.Value,
		Type:  t.Attribute.Type,
	}
	return CategoryRepo.CreateAttribute(categoryAttribute)
}
