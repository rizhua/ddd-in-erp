package entity

import (
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var BrandRepo repository.Brand

// 品牌
type Brand struct {
	ID   int64
	Name string
	Logo string
}

func (t *Brand) Create(orgID int64) error {
	brand := po.Brand{
		Name: t.Name,
		Logo: t.Logo,
	}
	return BrandRepo.Create(brand)
}

func (t *Brand) Delete(id []int64) error {
	return BrandRepo.Delete(id)
}

func (t *Brand) Update(orgID int64) error {
	brand := po.Brand{
		ID:   t.ID,
		Name: t.Name,
		Logo: t.Logo,
	}
	return BrandRepo.Update(brand)
}
