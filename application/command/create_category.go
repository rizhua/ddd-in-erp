package command

type CreateCategory struct {
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}
