package command

type CreateBrand struct {
	Name   string `json:"name" binding:"required"`
	Logo   string `json:"logo"`
	Letter string `json:"letter"`
	OrgID  int64  `json:"orgId"`
}
