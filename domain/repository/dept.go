package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 部门
type Dept interface {
	Create(po.Dept) error
	Update(po.Dept) error
	Delete(orgID int64, id []int64) error
	Find(orgID int64, req query.Request) ([]po.Dept, error)
}
