package application

import (
	"context"
	"encoding/json"
	"errors"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/converter"
)

func NewStructureService(
	structureDomain domain.StructureService,
	orgRepo repository.Org,
	deptRepo repository.Dept,
	roleRepo repository.Role,
	empRepo repository.Emp,
	userDomain domain.UserService,
) StructureService {
	return StructureService{
		structureDomain: structureDomain,
		orgRepo:         orgRepo,
		deptRepo:        deptRepo,
		roleRepo:        roleRepo,
		empRepo:         empRepo,
		userDomain:      userDomain,
	}
}

type StructureService struct {
	structureDomain domain.StructureService
	orgRepo         repository.Org
	deptRepo        repository.Dept
	roleRepo        repository.Role
	empRepo         repository.Emp
	userDomain      domain.UserService
	Context         context.Context
}

// 组织列表
func (t *StructureService) Find(args []byte) (any, error) {
	req := query.NewRequest()
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.orgRepo.Find(req)
	if err != nil {
		return nil, err
	}

	res := make(map[string]any)
	res["list"] = new(converter.Org).List(ret)
	res["total"] = cnt

	return res, err
}

func (t *StructureService) CreateDept(args []byte) error {
	cmd := command.CreateDept{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	t.structureDomain.Context = t.Context
	return t.structureDomain.CreateDept(cmd)
}

func (t *StructureService) DeleteDept(args []byte) error {
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

	return t.deptRepo.Delete(user.Org.ID, cmd.ID)
}

func (t *StructureService) UpdateDept(args []byte) error {
	cmd := command.UpdateDept{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	cmd.OrgID = user.Org.ID
	if cmd.OrgID == 0 {
		return errors.New("未取得组织信息")
	}
	return t.structureDomain.UpdateDept(cmd)
}

func (t *StructureService) FindDept(args []byte) (any, error) {
	req := query.NewRequest()
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return nil, err
	}
	ret, err := t.deptRepo.Find(user.Org.ID, req)
	if err != nil {
		return nil, err
	}
	return new(converter.Dept).Tree(ret, 0), err
}

func (t *StructureService) CreateEmp(args []byte) error {
	cmd := command.CreateEmp{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	cmd.OrgID = user.Org.ID
	if cmd.OrgID == 0 {
		return errors.New("未取得组织信息")
	}
	return t.structureDomain.CreateEmp(cmd)
}

func (t *StructureService) UpdateEmp(args []byte) error {
	cmd := command.UpdateEmp{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.userDomain.Context = t.Context
	user, err := t.userDomain.Parse()
	if err != nil {
		return err
	}
	cmd.OrgID = user.Org.ID
	if cmd.OrgID == 0 {
		return errors.New("未取得组织信息")
	}
	return t.structureDomain.UpdateEmp(cmd)
}

func (t *StructureService) FindEmp(args []byte) (any, error) {
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
	ret, cnt, err := t.empRepo.Find(req)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	data["list"] = new(converter.Emp).List(ret)
	data["total"] = cnt

	return data, err
}

func (t *StructureService) FindNode() (any, error) {
	t.structureDomain.Context = t.Context
	ret, err := t.structureDomain.FindNode()
	data := new(converter.Node).Tree(ret, 0)
	return data, err
}

func (t *StructureService) Switch(args []byte) error {
	cmd := command.UpdateOrg{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.structureDomain.Context = t.Context
	return t.structureDomain.Switch(cmd.ID)
}
