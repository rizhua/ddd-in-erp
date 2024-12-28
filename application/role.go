package application

import (
	"context"
	"encoding/json"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/converter"
)

func NewRoleService(
	roleDomain domain.RoleService,
	roleRepo repository.Role,
	userDomain domain.UserService,
) RoleService {
	return RoleService{
		roleDomain: roleDomain,
		roleRepo:   roleRepo,
		userDomain: userDomain,
	}
}

type RoleService struct {
	Context    context.Context
	roleDomain domain.RoleService
	roleRepo   repository.Role
	userDomain domain.UserService
}

func (t *RoleService) Create(args []byte) error {
	cmd := command.CreateRole{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	cmd.OrgID = user.Org.ID
	return t.roleDomain.Create(cmd)
}

func (t *RoleService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	return t.roleDomain.Delete(user.Org.ID, cmd.ID)
}

func (t *RoleService) Update(args []byte) error {
	cmd := command.UpdateRole{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	cmd.OrgID = user.Org.ID
	return t.roleDomain.Update(cmd)
}

func (t *RoleService) Find(args []byte) (any, error) {
	req := query.NewRequest(1, 1000)
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return nil, err
	}
	req.QueryBy = append(req.QueryBy, query.QueryBy{
		Field: "orgId",
		Value: user.Org.ID,
	})
	ret, err := t.roleRepo.Find(req)
	if err != nil {
		return nil, err
	}

	return new(converter.Role).Tree(ret, 0), nil
}

func (t *RoleService) BindNodeID(args []byte) error {
	cmd := command.BindingNode{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.roleRepo.BindNodeID(cmd.RoleID, cmd.NodeID)
}

func (t *RoleService) FindNodeID(args []byte) (any, error) {
	req := query.RoleBinding{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, err := t.roleRepo.FindNodeID(req.RoleID)
	if err != nil {
		return nil, err
	}
	var res []int64
	for _, v := range ret {
		res = append(res, v.NodeID)
	}

	return res, err
}

func (t *RoleService) FindNode(args []byte) (any, error) {
	req := query.RoleBinding{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, err := t.roleRepo.FindNode(req.RoleID, req.Request)
	if err != nil {
		return nil, err
	}
	res := new(converter.Node).Tree(ret, 0)

	return res, err
}

func (t *RoleService) AddUser(args []byte) error {
	cmd := command.BindingUser{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.roleRepo.AddUser(cmd.RoleID, cmd.UserID)
}

func (t *RoleService) RemoveUser(args []byte) error {
	cmd := command.BindingUser{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.roleRepo.RemoveUser(cmd.RoleID, cmd.UserID)
}

func (t *RoleService) FindUser(args []byte) (any, error) {
	req := query.RoleBinding{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return nil, err
	}
	ret, total, err := t.roleRepo.FindUser(user.Org.ID, req.RoleID, req.Request)
	if err != nil {
		return nil, err
	}
	res := new(converter.Emp).List(ret)
	data := map[string]interface{}{
		"list":  res,
		"total": total,
	}

	return data, err
}
