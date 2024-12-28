package query

type RoleBinding struct {
	RoleID int64 `json:"roleID" binding:"required"`
	Request
}
