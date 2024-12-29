package command

// 创建品牌
type CreateBrand struct {
	Name   string `json:"name" binding:"required"`
	Logo   string `json:"logo"`
	Letter string `json:"letter"`
	OrgID  int64  `json:"orgId"`
}

// 更新品牌
type UpdateBrand struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Letter string `json:"letter"`
	Sort   int8   `json:"sort"`
}
