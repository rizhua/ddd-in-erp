package entity

import (
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var DeptRepo repository.Dept

// 根实体:部门
type Dept struct {
	ID       int64
	Name     string
	ParentID int64
	Mgr      *Emp
	Org      *Org
}

func (t *Dept) Create() error {
	dept := po.Dept{
		Name:     t.Name,
		ParentID: t.ParentID,
		MgrID:    t.Mgr.ID,
		OrgID:    t.Org.ID,
	}
	return DeptRepo.Create(dept)
}

func (t *Dept) Update() error {
	dept := po.Dept{
		ID:       t.ID,
		Name:     t.Name,
		ParentID: t.ParentID,
		MgrID:    t.Mgr.ID,
		OrgID:    t.Org.ID,
	}
	return DeptRepo.Update(dept)
}
