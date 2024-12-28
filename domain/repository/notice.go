package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 公告
type Notice interface {
	Create(po.Notice) error
	Delete(id []int64) error
	Update(po.Notice) error
	Find(query.Request) ([]po.Notice, int64, error)
}
