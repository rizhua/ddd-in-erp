package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Spu interface {
	Create(po.Spu) error
	Delete(id []int64) error
	Update(po.Spu) error
	Get(id int64) (po.Spu, error)
	Find(query.Request) ([]po.Spu, int64, error)
}
