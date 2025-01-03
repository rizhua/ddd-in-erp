package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 商品类目
type Category interface {
	// 创建类目
	Create(po.Category) error

	// 删除类目
	Delete(id []int64) error

	// 更新类目
	Update(po.Category) error

	// 查找类目
	Find(query.Request) ([]po.Category, int64, error)

	// 设置类目排序
	SetSort(id int64, sort int8) error

	// 创建类目属性
	CreateAttribute(po.CategoryAttribute) error

	// 更新类目属性
	UpdateAttribute(po.CategoryAttribute) error

	// 删除类目属性
	DeleteAttribute(categoryID int64, id []int64) error

	// 查找类目属性
	FindAttribute(req query.Request) ([]po.CategoryAttribute, int64, error)
}
