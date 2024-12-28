package converter

import "rizhua.com/infrastructure/persistence/po"

type Bundle struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Term     int    `json:"term"`
	Price    int32  `json:"price"`
	Quota    int32  `json:"quota"`
	UpdateAt string `json:"updateAt"`
	CreateAt string `json:"createAt"`
}

func (t *Bundle) Info(in po.Bundle) (info Bundle) {
	info = Bundle{
		ID:       in.ID,
		Name:     in.Name,
		Term:     in.Term,
		Price:    in.Price,
		Quota:    in.Quota,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Bundle) List(in []po.Bundle) (list []Bundle) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
