package command

import "errors"

type CreateBundle struct {
	Name     string `json:"name" binding:"required"`
	Term     int    `json:"term"`
	Quota    int32  `json:"quota"`
	Price    int32  `json:"price"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *CreateBundle) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
