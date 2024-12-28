package query

import (
	"errors"

	"rizhua.com/pkg/util"
)

type ForgetUser struct {
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Captcha string `json:"captcha"`
}

func (t *ForgetUser) Validate() error {
	if t.Email != "" && util.IsEmail(t.Email) {
		return nil
	}
	if t.Mobile != "" && util.IsMobile(t.Mobile) {
		if t.Captcha == "" {
			return errors.New("验证码不能为空")
		}
		return nil
	}

	return errors.New("账号格式不正确，只能是手机号或邮箱")
}
