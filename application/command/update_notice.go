package command

type UpdateNotice struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Attach  string `json:"attach"`
	Scope   int8   `json:"scope"`
	Drafter string `json:"drafter"`
	Type    int8   `json:"type"`
}
