package command

type UpdateCategory struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentID int64  `json:"parentId"`
	Sort     int8   `json:"sort"`
}
