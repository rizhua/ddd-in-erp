package command

type UpdateAddress struct {
	ID      int64  `json:"id" binding:"required"`
	Contact string `json:"contact"`
	Tel     string `json:"tel"`
	Region  string `json:"region"`
	Detail  string `json:"detail"`
	Tag     string `json:"tag"`
	Default bool   `json:"default"`
}

type UpdateAddressDefault struct {
	ID int64 `json:"id" binding:"required"`
}
