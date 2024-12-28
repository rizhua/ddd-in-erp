package domain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	timeconv "github.com/Andrew-M-C/go.timeconv"
	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/adapter"
	"rizhua.com/infrastructure/constant"
	"rizhua.com/infrastructure/persistence/po"
)

// 组织架构
func NewStructureService(
	orgRepo repository.Org,
	deptRepo repository.Dept,
	userRepo repository.User,
	bundleRepo repository.Bundle,
	nodeRepo repository.Node,
) StructureService {
	entity.OrgRepo = orgRepo
	entity.DeptRepo = deptRepo
	entity.UserRepo = userRepo
	entity.BundleRepo = bundleRepo
	entity.NodeRepo = nodeRepo
	return StructureService{
		orgRepo:    orgRepo,
		deptRepo:   deptRepo,
		userRepo:   userRepo,
		bundleRepo: bundleRepo,
		nodeRepo:   nodeRepo,
	}
}

type StructureService struct {
	Context    context.Context
	orgRepo    repository.Org
	deptRepo   repository.Dept
	userRepo   repository.User
	bundleRepo repository.Bundle
	nodeRepo   repository.Node
}

// 创建部门
func (t *StructureService) CreateDept(cmd command.CreateDept) error {
	user, err := new(entity.User).Parse(t.Context)
	if err != nil {
		return err
	}
	if user.Org == nil || user.Org.ID == 0 {
		return errors.New("未获取到当前用户的组织，无法创建部门")
	}
	dept := entity.Dept{
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
		Mgr:      &entity.Emp{ID: cmd.MgrID},
		Org:      user.Org,
	}
	return dept.Create()
}

// 更新部门
func (t *StructureService) UpdateDept(cmd command.UpdateDept) error {
	user, err := new(entity.User).Parse(t.Context)
	if err != nil {
		return err
	}
	dept := entity.Dept{
		ID:       cmd.ID,
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
		Mgr:      &entity.Emp{ID: cmd.MgrID},
		Org:      user.Org,
	}
	return dept.Update()
}

// 创建员工
func (t *StructureService) CreateEmp(cmd command.CreateEmp) error {
	// 根据手机号获取用户信息
	userID, err := t.userRepo.GetByAccount(cmd.Mobile)
	if err != nil {
		return err
	}
	emp := entity.Emp{
		Name:     cmd.Name,
		Number:   cmd.Number,
		Grade:    cmd.Grade,
		Gender:   cmd.Gender,
		Position: cmd.Position,
		Tel:      cmd.Tel,
		Email:    cmd.Email,
		Address:  cmd.Address,
		JoinTime: cmd.JoinTime,
	}
	return emp.Create(cmd.OrgID, userID)
}

// 更新员工
func (t *StructureService) UpdateEmp(cmd command.UpdateEmp) error {
	emp := entity.Emp{
		ID:       cmd.ID,
		Name:     cmd.Name,
		Number:   cmd.Number,
		Grade:    cmd.Grade,
		Gender:   cmd.Gender,
		Position: cmd.Position,
		Tel:      cmd.Tel,
		Email:    cmd.Email,
		Address:  cmd.Address,
		JoinTime: cmd.JoinTime,
	}
	return emp.Update(cmd.OrgID)
}

// 组织权限
func (t *StructureService) FindNode() (list []po.Node, err error) {
	// 1、通过license 获取资源包信息
	userDO := entity.User{}
	user, err := userDO.Parse(t.Context)
	if err != nil {
		return
	}
	bundleDO := entity.Bundle{}
	bundle, err := bundleDO.GetByCode(user.Org.License)
	if err != nil {
		return
	}
	// 2、如果为付费资源包,生成的license只能单组织使用
	if bundle.Price > 0 {
		cnt, _ := t.orgRepo.Count(query.Request{
			QueryBy: []query.QueryBy{
				{
					Field: "license",
					Value: user.Org.License,
				},
			},
		})
		if cnt > 0 {
			err = errors.New("该授权码已有组织使用，请更换授权码")
			return
		}
		// 3、是否超过期限
		if bundle.Term > 0 {
			// 计算失效时间
			now := time.Now()
			expireAt := timeconv.AddDate(bundle.License.CreateAt, 0, bundle.Term, 0)
			if expireAt.Before(now) {
				err = errors.New("授权码已过期")
				return
			}
		}
	}
	list, err = t.bundleRepo.FindNode(bundle.ID, "")
	return
}

// 切换组织
func (t *StructureService) Switch(orgID int64) error {
	user, err := new(entity.User).Parse(t.Context)
	if err != nil {
		return err
	}
	// 根据 orgID 查询组织信息
	ret, err := new(entity.Org).FindByUserID(user.ID)
	if err != nil {
		return err
	}
	for _, v := range ret {
		if v.ID == orgID {
			user.Org = &v
			break
		}
	}
	// 根据 userID,orgID 查询员工
	if user.Org == nil {
		return errors.New("组织不存在")
	}
	emp := entity.Emp{}
	user.Emp, err = emp.Get(user.Org.ID, user.ID)
	if err != nil || user.Emp == nil {
		return errors.New("员工不存在")
	}
	// 缓存登录信息
	cache := adapter.NewCache()
	key := fmt.Sprintf("%v", constant.USER)
	field := fmt.Sprintf("%d", user.ID)
	value, _ := json.Marshal(user)
	err = cache.HSet(key, field, value)

	return err
}
