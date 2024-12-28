package converter

import "rizhua.com/infrastructure/persistence/po"

type Dept struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	MgrID    int64  `json:"mgrId"`
	// EmpCount int64  `json:"empCount"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Dept) Info(in po.Dept) (info Dept) {
	info = Dept{
		ID:       in.ID,
		Name:     in.Name,
		ParentID: in.ParentID,
		MgrID:    in.MgrID,
		// EmpCount: in.EmpCount,
		// UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		// CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Dept) List(in []po.Dept, parentID int64) (list []Dept) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}

type DeptTree struct {
	Dept
	Leaf     bool       `json:"leaf"`
	Children []DeptTree `json:"children"`
}

func (t *Dept) Tree(in []po.Dept, parentID int64) (list []DeptTree) {
	for _, v := range in {
		if parentID == v.ParentID {
			temp := DeptTree{
				Dept: t.Info(v),
				Leaf: false,
			}

			child := t.Tree(in, v.ID)
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
