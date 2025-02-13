package po

import "time"

// 标准化产品单元
type Spu struct {
	ID         int64  `xorm:"id pk autoincr"`
	Code       string `xorm:"code varchar(32)"`
	Name       string `xorm:"name varchar(240)"`
	LowPrice   int32  `xorm:"low_price int"`
	CategoryID int64  `xorm:"category_id int"`
	BrandID    int64  `xorm:"brand_id int"`
	SaleCount  int32  `xorm:"sale_count int"`
	RateCount  int32  `xorm:"rate_count int"`
	OrgID      int64  `xorm:"org_id int"`
	Barcode    string `xorm:"barcode text"`
	// Media      []map[string]any `xorm:"media json"`
	Detail   string    `xorm:"detail text"`
	Status   int8      `xorm:"status int"`
	UpdateAt time.Time `xorm:"update_at datetime"`
	CreateAt time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Spu) TableName() string {
	return "spu"
}

func (t *Spu) BeforeInsert() {
	t.Status = 1
	dt := time.Now()
	t.CreateAt = dt
	t.UpdateAt = dt
}

func (t *Spu) BeforeUpdate() {
	t.UpdateAt = time.Now()
}
