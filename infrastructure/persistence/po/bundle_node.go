package po

// 套餐节点
type BundleNode struct {
	BundleID int64 `xorm:"bundle_id bigint"`
	NodeID   int64 `xorm:"node_id bigint"`
}

// 设定表名
func (t *BundleNode) TableName() string {
	return "bundle_node"
}
