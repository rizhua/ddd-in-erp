package command

type CreateAddress struct {
	Contact string `json:"contact" binding:"required"`
	Tel     string `json:"tel"`
	Region  string `json:"region" binding:"required"`
	Detail  string `json:"detail" binding:"required"`
	Tag     string `json:"tag"`
	Default bool   `json:"default"`
}
