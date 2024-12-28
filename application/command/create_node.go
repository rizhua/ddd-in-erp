package command

import "errors"

type CreateNode struct {
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Meta     string `json:"meta"`
	Type     int8   `json:"type"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}

func (t *CreateNode) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
