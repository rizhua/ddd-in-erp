package command

type UpdateBrand struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Letter string `json:"letter"`
	Sort   int8   `json:"sort"`
}
