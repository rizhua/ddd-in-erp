package command

import "time"

type UpdateEmp struct {
	ID       int64     `json:"id" binding:"required"`
	Name     string    `json:"name"`
	Number   string    `json:"number"`
	Gender   int8      `json:"gender"`
	Position string    `json:"position"`
	Grade    string    `json:"grade"`
	Tel      string    `json:"tel"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
	JoinTime time.Time `json:"joinTime"`
	OrgID    int64
}
