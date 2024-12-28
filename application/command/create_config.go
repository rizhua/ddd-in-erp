package command

type CreateConfig struct {
	Code   string                 `json:"code"`
	Data   map[string]interface{} `json:"data"`
	Remark string
}
