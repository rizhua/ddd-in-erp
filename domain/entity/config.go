package entity

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var ConfigRepo repository.Config

type ConfigEntity struct {
	ID     int64
	Code   string
	Data   map[string]interface{}
	Remark string
}

func (t *ConfigEntity) Create(cmd command.CreateConfig) error {
	po := po.Config{
		Code:   cmd.Code,
		Data:   cmd.Data,
		Remark: cmd.Remark,
	}
	return ConfigRepo.Create(po)
}

func (t *ConfigEntity) Update(cmd command.UpdateConfig) error {
	po := po.Config{
		ID:   cmd.ID,
		Code: cmd.Code,
		Data: cmd.Data,
	}
	return ConfigRepo.Update(po)
}

func (t *ConfigEntity) GetByCode(code string) (po.Config, error) {
	po := po.Config{
		Code: code,
	}
	return ConfigRepo.Get(po)
}
