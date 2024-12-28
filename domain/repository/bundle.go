package repository

import (
	"rizhua.com/application/query"
	"rizhua.com/infrastructure/persistence/po"
)

type Bundle interface {
	// 创建套餐
	Create(po.Bundle) error

	// 删除套餐
	Delete(id []int64) error

	// 更新套餐
	Update(po.Bundle) error

	// 套餐详情
	Get(query.Request) (po.Bundle, error)

	// 套餐列表
	Find(query.Request) ([]po.Bundle, int64, error)

	// 套餐节点 id 列表
	FindNodeID(bundleID int64) ([]po.BundleNode, error)

	// 绑定、解绑节点
	BindNodeID(bundleID int64, nodeID []int64) error

	// 套餐节点列表
	FindNode(bundleID int64, path string) ([]po.Node, error)

	// 许可详情
	GetLicense(query.Request) (po.License, error)

	// 许可列表
	FindLicense(query.Request) ([]po.License, int64, error)
}
