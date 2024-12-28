package entity

import (
	"time"

	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var EmpRepo repository.Emp

// 根实体：员工
type Emp struct {
	ID       int64
	Name     string
	Number   string
	Gender   int8
	Position string
	Grade    string
	Tel      string
	Email    string
	Address  string
	JoinTime time.Time
	QuitTime time.Time
}

func (t *Emp) Create(orgID, userID int64) error {
	emp := po.Emp{
		Name:     t.Name,
		Number:   t.Number,
		Grade:    t.Grade,
		Gender:   t.Gender,
		Position: t.Position,
		Tel:      t.Tel,
		Email:    t.Email,
		Address:  t.Address,
		JoinTime: t.JoinTime,
		OrgID:    orgID,
		UserID:   userID,
	}
	return EmpRepo.Create(emp)
}

func (t *Emp) Update(orgID int64) error {
	emp := po.Emp{
		ID:       t.ID,
		Name:     t.Name,
		Number:   t.Number,
		Grade:    t.Grade,
		Gender:   t.Gender,
		Position: t.Position,
		Tel:      t.Tel,
		Email:    t.Email,
		Address:  t.Address,
		JoinTime: t.JoinTime,
		OrgID:    orgID,
	}
	return EmpRepo.Update(emp)
}

func (t *Emp) toEntity(in po.Emp) *Emp {
	return &Emp{
		ID:       in.ID,
		Name:     in.Name,
		Number:   in.Number,
		Grade:    in.Grade,
		Gender:   in.Gender,
		Position: in.Position,
		Tel:      in.Tel,
		Email:    in.Email,
		Address:  in.Address,
		JoinTime: in.JoinTime,
	}
}

func (t *Emp) Get(orgID, userID int64) (*Emp, error) {
	emp := po.Emp{
		UserID: userID,
		OrgID:  orgID,
	}
	ret, err := EmpRepo.Get(emp)
	if err != nil {
		return nil, err
	}
	return t.toEntity(ret), err
}

func (t *Emp) FindByUserID(userID int64) (list []Emp, err error) {
	ret, err := EmpRepo.FindByUserID(userID)
	for _, v := range ret {
		list = append(list, Emp{
			ID:       v.ID,
			Name:     v.Name,
			Number:   v.Number,
			Grade:    v.Grade,
			Gender:   v.Gender,
			Position: v.Position,
			Tel:      v.Tel,
			Email:    v.Email,
			Address:  v.Address,
			JoinTime: v.JoinTime,
		})
	}
	return
}
