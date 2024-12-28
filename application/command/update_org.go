package command

type UpdateOrg struct {
	ID int64 `json:"id" binding:"required"`
}
