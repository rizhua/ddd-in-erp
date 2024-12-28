package assembler

import (
	"rizhua.com/domain/entity"
)

type Dept struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	MgrID    int64  `json:"mgrId"`
}

func (t *Dept) Info(in entity.Dept) (info Dept) {
	info = Dept{
		ID:       in.ID,
		Name:     in.Name,
		ParentID: in.ParentID,
		MgrID:    in.Mgr.ID,
	}
	return
}

func (t *Dept) List(from []entity.Dept, parentID int64) (list []Dept) {
	for _, v := range from {
		list = append(list, t.Info(v))
	}
	return
}

type DeptTree struct {
	Dept
	Leaf     bool       `json:"leaf"`
	Children []DeptTree `json:"children"`
}

func (t *Dept) Tree(from []entity.Dept, parentID int64) (list []DeptTree) {
	for _, v := range from {
		if parentID == v.ParentID {
			temp := DeptTree{
				Dept: t.Info(v),
				Leaf: false,
			}

			child := t.Tree(from, v.ID)
			if len(child) == 0 {
				temp.Leaf = true
				temp.Children = make([]DeptTree, 0)
			} else {
				temp.Leaf = false
				temp.Children = child
			}
			list = append(list, temp)
		}
	}
	return
}
