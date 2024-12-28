package query

import (
	"errors"
)

type LoginUser struct {
	Account  string `json:"account"`
	Mode     int8   `json:"mode"` // 0: 密码, 1: 验证码
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

func (u *LoginUser) Validate() error {
	if u.Account == "" {
		return errors.New("账号不能为空")
	}

	if u.Mode == 0 && u.Password == "" {
		return errors.New("密码不能为空")
	}

	if u.Mode == 1 && u.Captcha == "" {
		return errors.New("验证码不能为空")
	}

	return nil
}
