package domain

import (
	"context"
	"errors"
	"time"

	timeconv "github.com/Andrew-M-C/go.timeconv"
	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

func NewNodeService(
	nodeRepo repository.Node,
	userRepo repository.User,
	bundleRepo repository.Bundle,
) NodeService {
	entity.NodeRepo = nodeRepo
	entity.UserRepo = userRepo
	entity.BundleRepo = bundleRepo
	return NodeService{
		nodeRepo:   nodeRepo,
		userRepo:   userRepo,
		bundleRepo: bundleRepo,
	}
}

type NodeService struct {
	nodeRepo   repository.Node
	userRepo   repository.User
	bundleRepo repository.Bundle
	Context    context.Context
}

func (t *NodeService) Create(cmd command.CreateNode) error {
	node := entity.Node{
		Icon:     cmd.Icon,
		Meta:     cmd.Meta,
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
		Sort:     cmd.Sort,
		Type:     cmd.Type,
	}
	return node.Create()
}

func (t *NodeService) Update(cmd command.UpdateNode) error {
	node := entity.Node{
		ID:   cmd.ID,
		Icon: cmd.Icon,
		Meta: cmd.Meta,
		Name: cmd.Name,
		Type: cmd.Type,
	}
	return node.Update()
}

func (t *NodeService) Permission(path string) (list []po.Node, err error) {
	userDO := entity.User{}
	user, err := userDO.Parse(t.Context)
	if err != nil {
		return
	}
	if user.Org != nil && user.Org.OwnerID == user.ID {
		// 组织的拥有者拥有该组织所有权限
		bundleDO := entity.Bundle{}
		bundle, err1 := bundleDO.GetByCode(user.Org.License)
		if err1 != nil {
			err = err1
			return
		}
		// 计算失效时间
		if bundle.Price > 0 && bundle.Term > 0 {
			now := time.Now()
			expireAt := timeconv.AddDate(bundle.License.CreateAt, 0, bundle.Term, 0)
			if expireAt.Before(now) {
				err = errors.New("授权码已过期")
				return
			}
		}
		list, err = t.bundleRepo.FindNode(bundle.ID, path)
	} else {
		// 组织成员，根据角色获取权限
		list, err = t.nodeRepo.Permission(user.ID, path)
	}
	return
}
