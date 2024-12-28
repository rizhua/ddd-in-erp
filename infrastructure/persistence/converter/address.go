package converter

import "rizhua.com/infrastructure/persistence/po"

type Address struct {
	ID       int64  `json:"id"`
	Contact  string `json:"contact"`
	Tel      string `json:"tel"`
	Region   string `json:"region"`
	Detail   string `json:"detail"`
	Default  bool   `json:"default"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Address) Info(po po.Address) (info Address) {
	info = Address{
		ID:       po.ID,
		Contact:  po.Contact,
		Tel:      po.Tel,
		Region:   po.Region,
		Detail:   po.Detail,
		Default:  po.Default,
		UpdateAt: po.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: po.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Address) List(po []po.Address) (list []Address) {
	for _, v := range po {
		list = append(list, t.Info(v))
	}
	return
}
