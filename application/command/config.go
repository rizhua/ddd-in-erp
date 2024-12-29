package command

type CreateConfig struct {
	Code   string                 `json:"code"`
	Data   map[string]interface{} `json:"data"`
	Remark string
}

type UpdateConfig struct {
	ID     int32                  `json:"id" binding:"required"`
	Code   string                 `json:"code"`
	Data   map[string]interface{} `json:"data"`
	Remark string                 `json:"remark"`
}
