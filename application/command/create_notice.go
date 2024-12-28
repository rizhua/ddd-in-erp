package command

type CreateNotice struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Attach      string `json:"attach"`
	Scope       int8   `json:"scope"`
	Drafter     string `json:"drafter"`
	DrafterDept string `json:"drafterDept"`
	Type        int8   `json:"type"`
}
