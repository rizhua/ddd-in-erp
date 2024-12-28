package entity

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var AddressRepo repository.Address

// 收货地址
type Address struct {
	ID      int64
	Contact string
	Detail  string
	Region  string
	Tag     string
	Tel     string
	Default bool
	UserID  int64
	OrgID   int64
}

func (a *Address) Create(cmd command.CreateAddress) error {
	addressPO := po.Address{
		Contact: cmd.Contact,
		Tel:     cmd.Tel,
		Region:  cmd.Region,
		Detail:  cmd.Detail,
		Tag:     cmd.Tag,
		Default: cmd.Default,
		UserID:  a.UserID,
		OrgID:   a.OrgID,
	}
	return AddressRepo.Create(addressPO)
}

func (a *Address) Update(cmd command.UpdateAddress) error {
	addressPO := po.Address{
		ID:      cmd.ID,
		Contact: cmd.Contact,
		Tel:     cmd.Tel,
		Region:  cmd.Region,
		Detail:  cmd.Detail,
		Tag:     cmd.Tag,
		Default: cmd.Default,
		UserID:  a.UserID,
		OrgID:   a.OrgID,
	}
	return AddressRepo.Update(addressPO)
}
