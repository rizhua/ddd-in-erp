package command

import (
	"errors"

	"rizhua.com/pkg/util"
)

type CreateUser struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha  string `json:"captcha"`
}

func (t *CreateUser) Validate() error {
	if util.IsEmail(t.Account) {
		return nil
	}

	if t.Captcha == "" {
		return errors.New("验证码不能为空")
	}
	return nil
}
