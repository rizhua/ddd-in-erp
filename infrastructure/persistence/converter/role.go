package converter

import "rizhua.com/infrastructure/persistence/po"

type Role struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	OrgID    int64  `json:"orgID"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Role) Info(in po.Role) (info Role) {
	info = Role{
		ID:       in.ID,
		Name:     in.Name,
		ParentID: in.ParentID,
		OrgID:    in.OrgID,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Role) List(in []po.Role) (list []Role) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}

type RoleTree struct {
	Role
	Leaf     bool       `json:"leaf"`
	Children []RoleTree `json:"children"`
}

func (t *Role) Tree(in []po.Role, parentID int64) (list []RoleTree) {
	for _, v := range in {
		if parentID == v.ParentID {
			temp := RoleTree{
				Role: t.Info(v),
				Leaf: false,
			}

			child := t.Tree(in, v.ID)
			if len(child) == 0 {
				temp.Leaf = true
				temp.Children = make([]RoleTree, 0)
			} else {
				temp.Leaf = false
				temp.Children = child
			}
			list = append(list, temp)
		}
	}
	return
}
