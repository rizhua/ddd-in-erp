package converter

import "rizhua.com/infrastructure/persistence/po"

type Attribute struct {
	ID       int64    `json:"id"`
	Label    string   `json:"label"`
	Value    []string `json:"value"`
	Multi    bool     `json:"multi"`
	Required bool     `json:"required"`
	UpdateAt string   `json:"updateAt"`
	CreateAt string   `json:"createAt"`
}

func (t *Attribute) Info(in po.Attribute) (info Attribute) {
	info = Attribute{
		ID:       in.ID,
		Label:    in.Label,
		Value:    in.Value,
		Multi:    in.Multi,
		Required: in.Required,
		UpdateAt: in.UpdateAt.Format("2006-01-02 15:04:05"),
		CreateAt: in.CreateAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (t *Attribute) List(in []po.Attribute) (list []Attribute) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
