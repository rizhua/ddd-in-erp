package entity

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var RoleRepo repository.Role

// 根实体：角色
type Role struct {
	ID       int64
	Name     string
	ParentID int64
	OrgID    int64
}

func (t *Role) Create(cmd command.CreateRole) error {
	rolePO := po.Role{
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
		OrgID:    t.OrgID,
	}
	return RoleRepo.Create(rolePO)
}

func (t *Role) Update(cmd command.UpdateRole) error {
	rolePO := po.Role{
		ID:       cmd.ID,
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
		OrgID:    t.OrgID,
	}
	return RoleRepo.Update(rolePO)
}
