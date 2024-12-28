package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Address interface {
	Create(po.Address) error
	Delete(id []int64, userID, orgID int64) error
	Update(po.Address) error
	Find(userID, OrgID int64, req query.Request) (list []po.Address, total int64, err error)
	SetDefault(id, orgID, userID int64) error
}
