package query

type Request struct {
	ID       int64     `json:"id"`
	Current  int       `json:"current"`
	PageSize int       `json:"pageSize" binding:"max=1000"`
	OrderBy  []OrderBy `json:"orderBy"`
	QueryBy  []QueryBy `json:"queryBy"`
	Omit     string    `json:"omit"`
	Keyword  string    `json:"keyword"`
}

func NewRequest(args ...int) Request {
	req := Request{
		Current:  1,
		PageSize: 10,
	}
	if len(args) > 0 {
		req.Current = args[0]
	}
	if len(args) > 1 {
		req.PageSize = args[1]
	}
	return req
}

type OrderBy struct {
	Field string `json:"field"`
	Value string `json:"value"` // asc desc
}

type QueryBy struct {
	Field string `json:"field"`
	Value any    `json:"value"`
}
