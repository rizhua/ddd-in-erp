package command

import "errors"

type UpdateBundle struct {
	ID    int64   `json:"id" binding:"required"`
	Name  string  `json:"name"`
	Term  int     `json:"term"`
	Quota int32   `json:"quota"`
	Price int32   `json:"price"`
	Node  []int64 `json:"node"`
}

func (t *UpdateBundle) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

type BindBundle struct {
	BundleID int64   `json:"bundleId" binding:"required"`
	NodeID   []int64 `json:"nodeId"`
}
