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

func NewBundleService(
	bundleDomain domain.BundleService,
	bundleRepo repository.Bundle,
	nodeRepo repository.Node,
	userDomain domain.UserService,
) BundleService {
	return BundleService{
		bundleDomain: bundleDomain,
		bundleRepo:   bundleRepo,
		nodeRepo:     nodeRepo,
		userDomain:   userDomain,
	}
}

type BundleService struct {
	Context      context.Context
	bundleDomain domain.BundleService
	bundleRepo   repository.Bundle
	nodeRepo     repository.Node
	userDomain   domain.UserService
}

func (t *BundleService) Create(args []byte) error {
	cmd := command.CreateBundle{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.bundleDomain.Create(cmd)
}

func (t *BundleService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	if err := cmd.Validate(); err != nil {
		return err
	}
	return t.bundleRepo.Delete(cmd.ID)
}

func (t *BundleService) Update(args []byte) error {
	cmd := command.UpdateBundle{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.bundleDomain.Update(cmd)
}

func (t *BundleService) Find(args []byte) (interface{}, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	data := make(map[string]any)
	ret, cnt, err := t.bundleRepo.Find(req)
	if err != nil {
		return data, err
	}

	data["list"] = new(converter.Bundle).List(ret)
	data["total"] = cnt

	return data, err
}

// 绑定、解绑节点
func (t *BundleService) BindNodeID(args []byte) error {
	cmd := command.BindBundle{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	err := t.bundleRepo.BindNodeID(cmd.BundleID, cmd.NodeID)
	if err != nil {
		return err
	}
	return nil
}

// 绑定的节点id列表
func (t *BundleService) FindNodeID(args []byte) (any, error) {
	req := query.Bundle{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, err := t.bundleRepo.FindNodeID(req.ID)
	if err != nil {
		return nil, err
	}
	list := make([]int64, len(ret))
	for i, v := range ret {
		list[i] = v.NodeID
	}
	return list, err
}

func (t *BundleService) FindLicense(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	data := make(map[string]any)
	ret, cnt, err := t.bundleRepo.FindLicense(req)
	if err != nil {
		return data, err
	}

	data["list"] = new(converter.License).List(ret)
	data["total"] = cnt

	return data, err
}
