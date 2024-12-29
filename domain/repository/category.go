package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 商品类目
// Category 接口定义了类目管理的相关方法
type Category interface {
	// Create 创建类目
	Create(po.Category) error

	// Delete 删除类目
	Delete(id []int64) error

	// Update 更新类目
	Update(po.Category) error

	// Find 查找类目
	Find(query.Request) ([]po.Category, int64, error)

	// SetSort 设置类目排序
	SetSort(id int64, sort int8) error

	// CreateAttribute 创建类目属性
	CreateAttribute(po.CategoryAttribute) error

	// UpdateAttribute 更新类目属性
	UpdateAttribute(po.CategoryAttribute) error

	// DeleteAttribute 删除类目属性
	DeleteAttribute(categoryID int64, id []int64) error

	// FindAttribute 查找类目属性
	FindAttribute(req query.Request) ([]po.CategoryAttribute, int64, error)
}
