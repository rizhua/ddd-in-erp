package command

// 创建地址
type CreateAddress struct {
	Contact string `json:"contact" binding:"required"`
	Tel     string `json:"tel"`
	Region  string `json:"region" binding:"required"`
	Detail  string `json:"detail" binding:"required"`
	Tag     string `json:"tag"`
	Default bool   `json:"default"`
}

// 更新地址
type UpdateAddress struct {
	ID      int64  `json:"id" binding:"required"`
	Contact string `json:"contact"`
	Tel     string `json:"tel"`
	Region  string `json:"region"`
	Detail  string `json:"detail"`
	Tag     string `json:"tag"`
	Default bool   `json:"default"`
}

// 设置默认
type UpdateAddressDefault struct {
	ID int64 `json:"id" binding:"required"`
}
