package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Role interface {
	Create(po.Role) error
	Update(po.Role) error
	Delete(orgID int64, id []int64) error
	GetByID(id int64) (po.Role, error)
	Find(req query.Request) ([]po.Role, error)

	BindNodeID(roleID int64, nodeID []int64) error
	FindNodeID(roleID int64) (list []po.RoleNode, err error)
	FindNode(roleID int64, req query.Request) (list []po.Node, err error)

	AddUser(roleID int64, userID []int64) error
	RemoveUser(roleID int64, userID []int64) error
	FindUser(orgID, roleID int64, req query.Request) (list []po.Emp, total int64, err error)
}
