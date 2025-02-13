package domain

import (
	"context"
	"fmt"

	"rizhua.com/application/command"
	"rizhua.com/application/query"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/adapter"
)

func NewProductService(
	spuRepo repository.Spu,
	categoryRepo repository.Category,
	attributeRepo repository.Attribute,
) ProductService {
	entity.SpuRepo = spuRepo
	entity.CategoryRepo = categoryRepo
	entity.AttributeRepo = attributeRepo
	return ProductService{
		spuRepo:       spuRepo,
		categoryRepo:  categoryRepo,
		attributeRepo: attributeRepo,
	}
}

type ProductService struct {
	spuRepo       repository.Spu
	categoryRepo  repository.Category
	attributeRepo repository.Attribute
	Context       context.Context
}

func (t *ProductService) Create(cmd command.CreateSpu) error {
	c := adapter.NewContext()
	fmt.Println(c)
	return nil
	// spuDO := entity.NewSpu(t.spuRepo)
	// return spuDO.Create(cmd)
}

func (t *ProductService) Delete(id []int64) error {
	return nil
}

func (t *ProductService) Update(cmd command.UpdateSpu) error {
	spuDO := entity.Spu{}
	return spuDO.Update(cmd)
}

func (t *ProductService) Get(id int64) (entity.Spu, error) {
	spuDO := entity.Spu{
		ID: id,
	}
	return spuDO.Get()
}

func (t *ProductService) Find(req query.Request) (list []entity.Spu, total int64, err error) {
	spu := entity.Spu{}
	return spu.Find(req)
}

func (t *ProductService) CreateAttribute(cmd command.CreateAttribute) error {
	attribute := entity.Attribute{
		Label:    cmd.Label,
		Value:    cmd.Value,
		Required: cmd.Required,
		Multi:    cmd.Multi,
	}
	return attribute.Create()
}

func (t *ProductService) UpdateAttribute(cmd command.UpdateAttribute) error {
	attribute := entity.Attribute{
		ID:       cmd.ID,
		Label:    cmd.Label,
		Value:    cmd.Value,
		Required: cmd.Required,
		Multi:    cmd.Multi,
	}
	return attribute.Update()
}
