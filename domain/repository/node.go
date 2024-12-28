package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Node interface {
	Create(po.Node) error
	Delete(id []int64) error
	Update(po.Node) error
	SetSort(po.Node) error
	SetStatus(po.Node) error
	GetByMeta(meta string) (po.Node, error)
	Find(query.Request) ([]po.Node, int64, error)

	Permission(userID int64, path string) ([]po.Node, error)
}
