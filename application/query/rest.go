package query

type Rest struct {
	Method string         `json:"method"`
	Params map[string]any `json:"params"`
}
