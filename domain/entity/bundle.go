package entity

import (
	"time"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var BundleRepo repository.Bundle

type license struct {
	ID       int64
	Code     string
	CreateAt time.Time
}

// 根实体:套餐
type Bundle struct {
	ID      int64
	Name    string
	Term    int
	Quota   int32
	Price   int32
	License *license
}

func (t *Bundle) Create(cmd command.CreateBundle) error {
	bundlePO := po.Bundle{
		Name:  cmd.Name,
		Price: cmd.Price,
		Quota: cmd.Quota,
		Term:  cmd.Term,
	}
	return BundleRepo.Create(bundlePO)
}

func (t *Bundle) Update(cmd command.UpdateBundle) error {
	bundlePO := po.Bundle{
		ID:    cmd.ID,
		Name:  cmd.Name,
		Price: cmd.Price,
		Quota: cmd.Quota,
		Term:  cmd.Term,
	}
	return BundleRepo.Update(bundlePO)
}

func (t *Bundle) Get(req query.Request) (info *Bundle, err error) {
	bundlePO, err := BundleRepo.Get(req)
	if err != nil {
		return
	}
	info = &Bundle{
		ID:    bundlePO.ID,
		Name:  bundlePO.Name,
		Price: bundlePO.Price,
		Quota: bundlePO.Quota,
		Term:  bundlePO.Term,
	}
	return
}

func (t *Bundle) GetByCode(code string) (*Bundle, error) {
	licensePO, err := BundleRepo.GetLicense(query.Request{
		QueryBy: []query.QueryBy{
			{
				Field: "code",
				Value: code,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	t.License = &license{
		ID:       licensePO.ID,
		Code:     licensePO.Code,
		CreateAt: licensePO.CreateAt,
	}
	bundle, err := BundleRepo.Get(query.Request{
		QueryBy: []query.QueryBy{
			{
				Field: "id",
				Value: licensePO.BizID,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &Bundle{
		ID:      bundle.ID,
		Name:    bundle.Name,
		Price:   bundle.Price,
		Quota:   bundle.Quota,
		Term:    bundle.Term,
		License: t.License,
	}, nil
}
