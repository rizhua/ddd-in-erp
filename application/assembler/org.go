package assembler

import (
	"rizhua.com/domain/entity"
)

type Org struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Contact  string `json:"contact"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
	License  string `json:"license"`
	OwnerID  int64  `json:"ownerId"`
	Status   int8   `json:"status"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Org) Info(org entity.Org) (info Org) {
	info = Org{
		ID:       org.ID,
		Name:     org.Name,
		FullName: org.FullName,
		Contact:  org.Contact,
		Tel:      org.Tel,
		Address:  org.Address,
		// OwnerID:  org.OwnerID,
		Status: org.Status,
		// UpdateAt: org.UpdateAt.Format("2006-01-02 15:04:05"),
		// CreateAt: org.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Org) List(org []entity.Org) (list []Org) {
	for _, v := range org {
		list = append(list, t.Info(v))
	}
	return
}
