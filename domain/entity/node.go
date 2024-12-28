package entity

import (
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var NodeRepo repository.Node

// 根实体：节点
type Node struct {
	ID       int64
	Icon     string
	Name     string
	Meta     string
	Type     int8
	ParentID int64
	Path     string
	Sort     int8
	Scheme   string
	Status   int8
}

func (t *Node) Create() error {
	node := po.Node{
		Icon:     t.Icon,
		Meta:     t.Meta,
		Name:     t.Name,
		ParentID: t.ParentID,
		Sort:     t.Sort,
		Type:     t.Type,
	}
	return NodeRepo.Create(node)
}

func (t *Node) Update() error {
	node := po.Node{
		ID:   t.ID,
		Icon: t.Icon,
		Meta: t.Meta,
		Name: t.Name,
		Type: t.Type,
	}
	return NodeRepo.Update(node)
}
