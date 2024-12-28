package assembler

import (
	"rizhua.com/domain/entity"
)

type Login struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Org      *Org   `json:"org"`
	Emp      *Emp   `json:"emp"`
}

func (t *Login) Info(in entity.User) Login {
	info := Login{
		ID:       in.ID,
		Nickname: in.Nickname,
		Email:    in.Email,
		Mobile:   in.Mobile,
	}
	if in.Org != nil {
		info.Org = &Org{
			ID:       in.Org.ID,
			Name:     in.Org.Name,
			FullName: in.Org.FullName,
			Tel:      in.Org.Tel,
			Address:  in.Org.Address,
			License:  in.Org.License,
			OwnerID:  in.Org.OwnerID,
			Status:   in.Org.Status,
		}
		info.Emp = &Emp{
			ID:       in.Emp.ID,
			Name:     in.Emp.Name,
			Number:   in.Emp.Number,
			Gender:   in.Emp.Gender,
			Grade:    in.Emp.Grade,
			Position: in.Emp.Position,
			Tel:      in.Emp.Tel,
			Email:    in.Emp.Email,
			Address:  in.Emp.Address,
			JoinTime: in.Emp.JoinTime.Format("2006-01-02 15:04:05"),
			QuitTime: in.Emp.QuitTime.Format("2006-01-02 15:04:05"),
		}
	}
	return info
}
