package command

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
