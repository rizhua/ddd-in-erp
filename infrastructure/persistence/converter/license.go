package converter

import "rizhua.com/infrastructure/persistence/po"

type License struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	BizID    int64  `json:"bizId"`
	CreateAt string `json:"createAt"`
}

func (t *License) Info(in po.License) (info License) {
	info = License{
		ID:       in.ID,
		Code:     in.Code,
		BizID:    in.BizID,
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *License) List(in []po.License) (list []License) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
