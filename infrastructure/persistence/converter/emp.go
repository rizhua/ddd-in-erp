package converter

import (
	"rizhua.com/infrastructure/persistence/po"
)

type Emp struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	Gender   int8   `json:"gender"`
	Position string `json:"position"`
	Grade    string `json:"grade"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Status   int8   `json:"status"`
	OrgID    int64  `json:"orgId"`
	JoinTime string `json:"joinTime"`
	QuitTime string `json:"quitTime"`
}

func (t *Emp) Info(in po.Emp) (info Emp) {
	info = Emp{
		ID:       in.ID,
		UserID:   in.UserID,
		Name:     in.Name,
		Number:   in.Number,
		Position: in.Position,
		Grade:    in.Grade,
		Tel:      in.Tel,
		Email:    in.Email,
		Address:  in.Address,
		Status:   in.Status,
		OrgID:    in.OrgID,
		JoinTime: in.JoinTime.Format("2006-01-02"),
		QuitTime: in.QuitTime.Format("2006-01-02"),
	}
	if info.JoinTime == "0001-01-01" {
		info.JoinTime = ""
	}
	if info.QuitTime == "0001-01-01" {
		info.QuitTime = ""
	}
	return
}

func (t *Emp) List(in []po.Emp) (list []Emp) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
