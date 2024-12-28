package command

type SendSms struct {
	Mobile string `json:"mobile" binding:"required"`
	Scene  int    `json:"scene"`
}
