package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 商品属性
type Attribute interface {
	// 创建属性
	Create(po.Attribute) error

	// 删除属性
	Delete(id []int64) error

	// 更新属性
	Update(po.Attribute) error

	// 属性列表
	Find(query.Request) ([]po.Attribute, int64, error)
}
