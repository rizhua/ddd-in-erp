package command

import "errors"

type Delete struct {
	ID []int64 `json:"id" binding:"required"`
}

func (t Delete) Validate() error {
	if len(t.ID) == 0 {
		return errors.New("参数错误")
	}

	return nil
}
