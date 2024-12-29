package command

// 创建商品
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

// 更新商品
type UpdateSpu struct {
	ID        int64                    `json:"id" binding:"required"`
	Code      string                   `json:"code"`
	Title     string                   `json:"title"`
	LowPrice  int32                    `json:"lowPrice"`
	RateCount int32                    `json:"rateCount"`
	Barcode   string                   `json:"barcode"`
	Media     []map[string]interface{} `json:"media"`
	Attribute []map[string]interface{} `json:"attribute"`
	Detail    string                   `json:"detail"`
	Status    int8                     `json:"status"`
}
