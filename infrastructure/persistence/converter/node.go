package converter

import "rizhua.com/infrastructure/persistence/po"

type Node struct {
	ID       int64  `json:"id"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Meta     string `json:"meta"`
	Type     int8   `json:"type"`
	ParentID int64  `json:"parentId"`
	Path     string `json:"path"`
	Sort     int8   `json:"sort"`
	Status   int8   `json:"status"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Node) Info(in po.Node) (info Node) {
	info = Node{
		ID:       in.ID,
		Icon:     in.Icon,
		Name:     in.Name,
		Meta:     in.Meta,
		Type:     in.Type,
		ParentID: in.ParentID,
		Path:     in.Path,
		Sort:     in.Sort,
		Status:   in.Status,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Node) List(in []po.Node) (list []Node) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}

type NodeTree struct {
	Node
	Leaf     bool       `json:"leaf"`
	Children []NodeTree `json:"children"`
}

func (t *Node) Tree(in []po.Node, parentID int64) (list []NodeTree) {
	for _, v := range in {
		if parentID == v.ParentID {
			temp := NodeTree{
				Node: t.Info(v),
				Leaf: false,
			}

			child := t.Tree(in, v.ID)
			if len(child) == 0 {
				temp.Leaf = true
				temp.Children = make([]NodeTree, 0)
			} else {
				temp.Leaf = false
				temp.Children = child
			}
			list = append(list, temp)
		}
	}
	return
}
