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

func NewCategoryService(
	categoryDomain domain.CategoryService,
	categoryRepo repository.Category,
) CategoryService {
	return CategoryService{
		categoryDomain: categoryDomain,
		categoryRepo:   categoryRepo,
	}
}

type CategoryService struct {
	categoryDomain domain.CategoryService
	categoryRepo   repository.Category
	Context        context.Context
}

func (t *CategoryService) Create(args []byte) error {
	cmd := command.CreateCategory{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.categoryDomain.Context = t.Context
	return t.categoryDomain.Create(cmd)
}

func (t *CategoryService) Delete(args []byte) error {
	cmd := command.Delete{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.categoryRepo.Delete(cmd.ID)
}

func (t *CategoryService) Update(args []byte) error {
	cmd := command.UpdateCategory{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.categoryDomain.Context = t.Context
	return t.categoryDomain.Update(cmd)
}

func (t *CategoryService) Find(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.categoryRepo.Find(req)
	if err != nil {
		return nil, err
	}
	res := new(converter.Category).Tree(ret, 0)
	data := map[string]any{
		"list":  res,
		"total": cnt,
	}

	return data, err
}

// 创建类目属性
func (t *CategoryService) CreateAttribute(args []byte) error {
	cmd := command.CreateCategoryAttribute{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.categoryDomain.Context = t.Context
	return t.categoryDomain.CreateAttribute(cmd)
}

// 更新类目属性
func (t *CategoryService) UpdateAttribute(args []byte) error {
	cmd := command.UpdateCategoryAttribute{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	t.categoryDomain.Context = t.Context
	return t.categoryDomain.UpdateAttribute(cmd)
}

// 删除类目属性
func (t *CategoryService) DeleteAttribute(args []byte) error {
	cmd := command.DeleteCategoryAttribute{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return err
	}
	return t.categoryRepo.DeleteAttribute(cmd.CategoryID, cmd.ID)
}

// 获取类目属性
func (t *CategoryService) FindAttribute(args []byte) (any, error) {
	req := query.Request{}
	if err := json.Unmarshal(args, &req); err != nil {
		return nil, err
	}
	ret, cnt, err := t.categoryRepo.FindAttribute(req)
	if err != nil {
		return nil, err
	}
	res := new(converter.CategoryAttribute).List(ret)
	data := map[string]any{
		"list":  res,
		"total": cnt,
	}

	return data, err
}
