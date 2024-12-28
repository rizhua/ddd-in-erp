package command

type CreateAttribute struct {
	Label    string   `json:"label" binding:"required"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
	IsSale   bool     `json:"isSale"`
}
