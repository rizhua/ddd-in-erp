package command

type CreateOrg struct {
	Name    string `json:"name" binding:"required"`
	Abbr    string `json:"abbr"`
	Contact string `json:"contact"`
	Tel     string `json:"tel"`
	Address string `json:"address"`
	OwnerID string `json:"ownerId"`
}

type UpdateOrg struct {
	ID int64 `json:"id" binding:"required"`
}
