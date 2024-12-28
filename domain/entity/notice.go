package entity

import (
	"rizhua.com/application/command"
	"rizhua.com/domain/repository"
	"rizhua.com/infrastructure/persistence/po"
)

var NoticeRepo repository.Notice

type NoticeEntity struct {
	ID      int64
	Title   string
	Content string
	Attach  string
	Type    int8
	Scope   int8
	Drafter string
}

func (t *NoticeEntity) Create(cmd command.CreateNotice) error {
	po := po.Notice{
		Title:   cmd.Title,
		Content: cmd.Content,
		Attach:  cmd.Attach,
		Type:    cmd.Type,
		Scope:   cmd.Scope,
		Drafter: cmd.Drafter,
	}
	return NoticeRepo.Create(po)
}

func (t *NoticeEntity) Update(cmd command.UpdateNotice) error {
	po := po.Notice{
		ID:      cmd.ID,
		Title:   cmd.Title,
		Content: cmd.Content,
		Attach:  cmd.Attach,
		Type:    cmd.Type,
		Scope:   cmd.Scope,
		Drafter: cmd.Drafter,
	}
	return NoticeRepo.Update(po)
}
