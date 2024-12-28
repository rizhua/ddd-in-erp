package repository

import "rizhua.com/infrastructure/persistence/po"

type Order interface {
	Create(po.Order) (int64, error)
}
