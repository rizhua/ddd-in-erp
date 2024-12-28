package command

import (
	"errors"
	"time"
)

type CreateEmp struct {
	Mobile   string    `json:"mobile" binding:"required"`
	Name     string    `json:"name"`
	Number   string    `json:"number"`
	Gender   int8      `json:"gender"`
	Position string    `json:"position"`
	Grade    string    `json:"grade"`
	Tel      string    `json:"tel"`
	Email    string    `json:"email"`
	MgrID    int64     `json:"mgrId"`
	Address  string    `json:"address"`
	JoinTime time.Time `json:"joinTime"`
	OrgID    int64
}

func (t *CreateEmp) Validate() error {
	if t.Mobile == "" {
		return errors.New("mobile is required")
	}
	return nil
}
