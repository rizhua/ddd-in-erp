package command

import "errors"

type Delete struct {
	ID []int64 `json:"id" binding:"required"`
}

func (t Delete) Validate() error {
	if len(t.ID) == 0 {
		return errors.New("参数错误")
	}

	return nil
}

// 发送邮件
type SendEmail struct {
	Email         string            `json:"email" binding:"required"`
	TemplateCode  string            `json:"templateCode"`
	TemplateParam map[string]string `json:"templateParam"`
}

// 发送短信
type SendSms struct {
	Mobile string `json:"mobile" binding:"required"`
	Scene  int    `json:"scene"`
}
