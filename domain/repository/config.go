package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 配置
type Config interface {
	Create(po.Config) error
	Delete(id []int64) error
	Update(po.Config) error
	Get(po.Config) (po.Config, error)
	Find(query.Request) ([]po.Config, int64, error)
}
