package assembler

import (
	"rizhua.com/domain/entity"
)

type Emp struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	Gender   int8   `json:"gender"`
	Position string `json:"position"`
	Grade    string `json:"grade"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Status   int8   `json:"status"`
	JoinTime string `json:"joinTime"`
	QuitTime string `json:"quitTime"`
}

func (t *Emp) Info(in entity.Emp) (info Emp) {
	info = Emp{
		ID:       in.ID,
		Name:     in.Name,
		Number:   in.Number,
		Gender:   in.Gender,
		Position: in.Position,
		Grade:    in.Grade,
		Tel:      in.Tel,
		Email:    in.Email,
		Address:  in.Address,
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

func (t *Emp) List(in []entity.Emp) (list []Emp) {
	for _, v := range in {
		list = append(list, t.Info(v))
	}
	return
}
