package command

import "errors"

type UpdateUser struct {
	Avatar   string `json:"avatar"`
	Birthday string `json:"birthday"`
	Gender   int8   `json:"gender"`
}

// 修改密码
type UpdateUserPassword struct {
	Code        string `form:"code"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword"`
}

func (t *UpdateUserPassword) Validate() error {
	if t.OldPassword == "" {
		return errors.New("原密码不能空")
	}
	return nil
}

// 忘记密码->重置密码
type ResetUserPassword struct {
	Token      string `json:"token" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword"`
}

func (t *ResetUserPassword) Validate() error {
	if t.Password != t.RePassword {
		return errors.New("两次输入的密码不一致")
	}
	return nil
}

// 激活帐号
type ActiveUser struct {
	Token string `json:"token" binding:"required"`
}
