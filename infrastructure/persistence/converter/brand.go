package converter

import (
	"rizhua.com/infrastructure/persistence/po"
)

type Brand struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Brand) Info(in po.Brand) (info Brand) {
	info = Brand{
		ID:       in.ID,
		Name:     in.Name,
		Logo:     in.Logo,
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Brand) List(in []po.Brand) (list []Brand) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
