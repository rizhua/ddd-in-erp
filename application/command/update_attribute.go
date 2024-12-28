package command

type UpdateAttribute struct {
	ID       int64    `json:"id" binding:"required"`
	Label    string   `json:"label"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
	IsSale   bool     `json:"isSale"`
}
