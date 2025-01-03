package command

import "github.com/go-playground/validator/v10"

// 创建属性
type CreateAttribute struct {
	Label    string   `json:"label" binding:"required"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
}

func (c *CreateAttribute) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

// 更新属性
type UpdateAttribute struct {
	ID       int64    `json:"id" binding:"required"`
	Label    string   `json:"label"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
}

func (c *UpdateAttribute) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
