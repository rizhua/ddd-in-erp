package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

// 用户
type User interface {
	// 创建用户
	Create(info po.User) (int64, error)

	// 更新用户
	Update(info po.User) error

	// 用户详情
	Get(info po.User) (po.User, error)

	// 根据账号获取用户ID
	GetByAccount(account string) (int64, error)

	// 用户列表
	Find(req query.Request) ([]po.User, int64, error)

	// 修改用户
	SetPassword(po.User) error

	// 用户是否存在
	Exist(po.User) bool
}
