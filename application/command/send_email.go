package command

type SendEmail struct {
	Email         string            `json:"email" binding:"required"`
	TemplateCode  string            `json:"templateCode"`
	TemplateParam map[string]string `json:"templateParam"`
}
