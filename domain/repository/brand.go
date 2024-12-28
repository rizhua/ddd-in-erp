package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 品牌
type Brand interface {
	// 创建品牌
	Create(po.Brand) error

	// 删除品牌
	Delete(id []int64) error

	// 更新品牌
	Update(po.Brand) error

	// 品牌列表
	Find(query.Request) ([]po.Brand, int64, error)
}
