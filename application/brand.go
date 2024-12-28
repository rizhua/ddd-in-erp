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

func NewBrandService(
	brandDomain domain.BrandService,
	brandRepo repository.Brand,
) BrandService {
	return BrandService{
		brandDomain: brandDomain,
		brandRepo:   brandRepo,
	}
}

type BrandService struct {
	Context     context.Context
	brandDomain domain.BrandService
	brandRepo   repository.Brand
}

func (t *BrandService) Create(args []byte) error {
	cmd := command.CreateBrand{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.brandDomain.Context = t.Context
	return t.brandDomain.Create(cmd)
}

func (t *BrandService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.brandDomain.Context = t.Context
	return t.brandRepo.Delete(cmd.ID)
}

func (t *BrandService) Update(args []byte) error {
	cmd := command.UpdateBrand{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.brandDomain.Context = t.Context
	return t.brandDomain.Update(cmd)
}

func (t *BrandService) Find(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.brandRepo.Find(req)
	if err != nil {
		return nil, err
	}
	res := new(converter.Brand).List(ret)
	data := map[string]any{
		"list":  res,
		"total": cnt,
	}

	return data, err
}
