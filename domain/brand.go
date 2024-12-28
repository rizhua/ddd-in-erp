package domain

import (
	"context"

	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
)

func NewBrandService(
	brandRepo repository.Brand,
) BrandService {
	entity.BrandRepo = brandRepo
	return BrandService{
		brandRepo: brandRepo,
	}
}

type BrandService struct {
	Context   context.Context
	brandRepo repository.Brand
}

func (t *BrandService) Create(cmd command.CreateBrand) error {
	user, err := new(entity.User).Parse(t.Context)
	if err != nil {
		return err
	}
	brand := entity.Brand{
		Name: cmd.Name,
		Logo: cmd.Logo,
	}
	return brand.Create(user.Org.ID)
}

func (t *BrandService) Update(cmd command.UpdateBrand) error {
	user, err := new(entity.User).Parse(t.Context)
	if err != nil {
		return err
	}
	brand := entity.Brand{
		ID:   cmd.ID,
		Name: cmd.Name,
		Logo: cmd.Logo,
	}
	return brand.Update(user.Org.ID)
}
