package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Emp interface {
	Create(po.Emp) error
	Update(po.Emp) error
	Get(po.Emp) (po.Emp, error)
	Find(req query.Request) (list []po.Emp, total int64, err error)
	FindByUserID(userID int64) (list []po.Emp, err error)
}
