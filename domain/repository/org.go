package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Org interface {
	Create(po.Org) error
	Update(po.Org) error
	GetByID(ID int64) (po.Org, error)
	Find(req query.Request) ([]po.Org, int64, error)
	FindByUserID(userID int64) ([]po.Org, error)
	Count(req query.Request) (int64, error)
}
