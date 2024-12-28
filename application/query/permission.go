package query

import "errors"

type Permission struct {
	Meta string `json:"meta"`
}

func (t Permission) Validate() error {
	if t.Meta == "" {
		return errors.New("参数错误: 缺少 meta 参数")
	}
	return nil
}
