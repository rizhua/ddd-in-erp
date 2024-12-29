package command

// 创建属性
type CreateAttribute struct {
	Label    string   `json:"label" binding:"required"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
	IsSale   bool     `json:"isSale"`
}

// 更新属性
type UpdateAttribute struct {
	ID       int64    `json:"id" binding:"required"`
	Label    string   `json:"label"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
	IsSale   bool     `json:"isSale"`
}
