package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Category interface {
	Create(po.Category) error
	Delete(id []int64) error
	Update(po.Category) error
	Find(query.Request) ([]po.Category, int64, error)
	SetSort(id int64, sort int8) error
	CreateAttribute(po.CategoryAttribute) error
	UpdateAttribute(po.CategoryAttribute) error
	DeleteAttribute(categoryID int64, id []int64) error
	FindAttribute(req query.Request) ([]po.CategoryAttribute, int64, error)
}
