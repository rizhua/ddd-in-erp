package entity

import (
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var AttributeRepo repository.Attribute

// 商品属性
type Attribute struct {
	ID       int64
	Label    string
	Value    []string
	Multi    bool
	Required bool
	IsSale   bool
}

func (t *Attribute) Create() error {
	attribute := po.Attribute{
		Label:    t.Label,
		Value:    t.Value,
		Multi:    t.Multi,
		Required: t.Required,
	}
	return AttributeRepo.Create(attribute)
}

func (t *Attribute) Update() error {
	attribute := po.Attribute{
		ID:       t.ID,
		Label:    t.Label,
		Value:    t.Value,
		Multi:    t.Multi,
		Required: t.Required,
	}
	return AttributeRepo.Update(attribute)
}
