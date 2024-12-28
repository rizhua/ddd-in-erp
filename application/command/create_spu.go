package command

type CreateSpu struct {
	Code       string           `json:"code"`
	Title      string           `json:"title" binding:"required"`
	LowPrice   int32            `json:"lowPrice"`
	Barcode    string           `json:"barcode"`
	Media      []map[string]any `json:"media"`
	CategoryID int64            `json:"categoryId"`
	BrandID    int64            `json:"brandId"`
	Attribute  []map[string]any `json:"attribute"`
	Detail     string           `json:"detail"`
	Status     int8             `json:"status"`
}
