package converter

import (
	"rizhua.com/infrastructure/persistence/po"
)

type Org struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Contact  string `json:"contact"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
	OwnerID  int64  `json:"ownerId"`
	Status   int8   `json:"status"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Org) Info(in po.Org) (info Org) {
	info = Org{
		ID:       in.ID,
		Name:     in.Name,
		FullName: in.FullName,
		Contact:  in.Contact,
		Tel:      in.Tel,
		Address:  in.Address,
		OwnerID:  in.OwnerID,
		Status:   in.Status,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Org) List(in []po.Org) (list []Org) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
