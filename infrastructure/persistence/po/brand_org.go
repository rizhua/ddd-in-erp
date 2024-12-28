package po

// 品牌代理
type BrandOrg struct {
	BrandID int64 `xorm:"brand_id bigint"`
	OrgID   int64 `xorm:"org_id bigint"`
}

// 设定表名
func (t *BrandOrg) TableName() string {
	return "brand_org"
}
