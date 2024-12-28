package persistence

import (
	"errors"

	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
	"xorm.io/xorm"
)

func NewOrder(db *xorm.Engine) repository.Order {
	return &order{db: db}
}

type order struct {
	db *xorm.Engine
}

// 创建订单
func (t *order) Create(info po.Order) (id int64, err error) {
	has, err := t.db.Insert(&info)
	if err != nil || has == 0 {
		err = errors.New("写入失败")
		return
	}
	id = info.ID
	return
}
