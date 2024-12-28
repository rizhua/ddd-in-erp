package entity

import (
	"rizhua.com/application/query"
	"rizhua.com/domain/repository"
)

var OrgRepo repository.Org

type OrgStatus int8

const (
	Disable OrgStatus = iota
	Enable
)

func (t OrgStatus) String() string {
	return [...]string{"禁用", "启用"}[t]
}

// 根实体：组织
type Org struct {
	ID       int64
	Icon     string
	Code     string
	Name     string
	FullName string
	Industry string
	Capacity int32
	Contact  string
	Tel      string
	Address  string
	OwnerID  int64
	License  string
	Status   int8
}

func (t *Org) GetByID() (*Org, error) {
	ret, err := OrgRepo.GetByID(t.ID)
	if err != nil {
		return nil, err
	}
	return &Org{
		ID:       ret.ID,
		Code:     ret.Code,
		Name:     ret.Name,
		FullName: ret.FullName,
		Icon:     ret.Icon,
		Contact:  ret.Contact,
		Tel:      ret.Tel,
		Address:  ret.Address,
		OwnerID:  ret.OwnerID,
		License:  ret.License,
		Status:   ret.Status,
	}, nil
}

func (t *Org) Find(req query.Request) (list []Org, total int64, err error) {
	ret, total, err := OrgRepo.Find(req)
	for _, v := range ret {
		list = append(list, Org{
			ID:       v.ID,
			Name:     v.Name,
			FullName: v.FullName,
			Contact:  v.Contact,
			Tel:      v.Tel,
			Address:  v.Address,
			Status:   v.Status,
		})
	}
	return
}

func (t *Org) FindByUserID(userID int64) (list []Org, err error) {
	ret, err := OrgRepo.FindByUserID(userID)
	for _, v := range ret {
		list = append(list, Org{
			ID:       v.ID,
			Code:     v.Code,
			Name:     v.Name,
			FullName: v.FullName,
			Icon:     v.Icon,
			Contact:  v.Contact,
			Tel:      v.Tel,
			Address:  v.Address,
			OwnerID:  v.OwnerID,
			License:  v.License,
			Status:   v.Status,
		})
	}
	return
}
