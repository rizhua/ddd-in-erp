package domain

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
)

func NewBundleService(
	bundleRepo repository.Bundle,
	nodeRepo repository.Node,
) BundleService {
	entity.BundleRepo = bundleRepo
	entity.NodeRepo = nodeRepo
	return BundleService{
		bundleRepo: bundleRepo,
		nodeRepo:   nodeRepo,
	}
}

type BundleService struct {
	bundleRepo repository.Bundle
	nodeRepo   repository.Node
}

func (t *BundleService) Create(cmd command.CreateBundle) error {
	bundleDO := entity.Bundle{}
	return bundleDO.Create(cmd)
}

func (t *BundleService) Update(cmd command.UpdateBundle) error {
	bundleDO := entity.Bundle{}
	return bundleDO.Update(cmd)
}
