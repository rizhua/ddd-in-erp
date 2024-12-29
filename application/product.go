package application

import (
	"context"
	"encoding/json"

	"rizhua.com/application/assembler"
	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/converter"
)

func NewProductService(
	productDomain domain.ProductService,
	spuRepo repository.Spu,
	categoryRepo repository.Category,
	attributeRepo repository.Attribute,
) ProductService {
	return ProductService{
		productDomain: productDomain,
		spuRepo:       spuRepo,
		categoryRepo:  categoryRepo,
		attributeRepo: attributeRepo,
	}
}

type ProductService struct {
	productDomain domain.ProductService
	spuRepo       repository.Spu
	categoryRepo  repository.Category
	attributeRepo repository.Attribute
	Context       context.Context
}

func (t *ProductService) Create(args []byte) error {
	cmd := command.CreateSpu{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.productDomain.Context = t.Context
	return t.productDomain.Create(cmd)
}

func (t *ProductService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.productDomain.Context = t.Context
	return t.productDomain.Delete(cmd.ID)
}

func (t *ProductService) Update(args []byte) error {
	cmd := command.UpdateSpu{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.productDomain.Context = t.Context
	return t.productDomain.Update(cmd)
}

func (t *ProductService) Get(args []byte) (any, error) {
	req := query.Product{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	ret, err := t.productDomain.Get(req.ID)
	return new(assembler.Spu).Info(ret), err
}

func (t *ProductService) Find(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	t.productDomain.Context = t.Context
	ret, cnt, err := t.productDomain.Find(req)
	if err != nil {
		return nil, err
	}
	res := new(assembler.Spu).List(ret)
	data := map[string]any{
		"list":  res,
		"total": cnt,
	}

	return data, err
}

func (t *ProductService) CreateAttribute(args []byte) error {
	cmd := command.CreateAttribute{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.productDomain.Context = t.Context
	return t.productDomain.CreateAttribute(cmd)
}

func (t *ProductService) DeleteAttribute(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.attributeRepo.Delete(cmd.ID)
}

func (t *ProductService) UpdateAttribute(args []byte) error {
	cmd := command.UpdateAttribute{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.productDomain.Context = t.Context
	return t.productDomain.UpdateAttribute(cmd)
}

func (t *ProductService) FindAttribute(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.attributeRepo.Find(req)
	if err != nil {
		return nil, err
	}
	res := new(converter.Attribute).List(ret)
	data := map[string]any{
		"list":  res,
		"total": cnt,
	}

	return data, err
}
