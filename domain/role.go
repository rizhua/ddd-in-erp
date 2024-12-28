package domain

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
)

func NewRoleService(
	roleRepo repository.Role,
) RoleService {
	return RoleService{
		roleRepo: roleRepo,
	}
}

type RoleService struct {
	roleRepo repository.Role
}

func (t *RoleService) Create(cmd command.CreateRole) error {
	roleDO := entity.Role{}
	return roleDO.Create(cmd)
}

func (t *RoleService) Delete(orgID int64, id []int64) error {
	return t.roleRepo.Delete(orgID, id)
}

func (t *RoleService) Update(cmd command.UpdateRole) error {
	roleDO := entity.Role{}
	return roleDO.Update(cmd)
}
