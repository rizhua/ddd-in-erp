package domain

import (
	"context"

	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
)

func NewCategoryService(
	categoryRepo repository.Category,
) CategoryService {
	entity.CategoryRepo = categoryRepo
	return CategoryService{
		categoryRepo: categoryRepo,
	}
}

type CategoryService struct {
	Context      context.Context
	categoryRepo repository.Category
}

func (t *CategoryService) Create(cmd command.CreateCategory) error {
	category := entity.Category{
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
	}
	return category.Create()
}

func (t *CategoryService) Update(cmd command.UpdateCategory) error {
	category := entity.Category{
		ID:       cmd.ID,
		Name:     cmd.Name,
		ParentID: cmd.ParentID,
	}
	return category.Update()
}

func (t *CategoryService) CreateAttribute(cmd command.CreateCategoryAttribute) error {
	category := entity.Category{
		ID: cmd.CategoryID,
		Attribute: &entity.CategoryAttribute{
			Label: cmd.Label,
			Value: cmd.Value,
			Type:  cmd.Type,
		},
	}
	return category.Create()
}

func (t *CategoryService) UpdateAttribute(cmd command.UpdateCategoryAttribute) error {
	category := entity.Category{
		Attribute: &entity.CategoryAttribute{
			ID:    cmd.ID,
			Label: cmd.Label,
			Value: cmd.Value,
			Type:  cmd.Type,
		},
	}
	return category.Update()
}
