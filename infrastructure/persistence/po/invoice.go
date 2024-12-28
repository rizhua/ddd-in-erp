package po

// 发票
type Invoice struct {
	SpuID     int32  `xorm:"pk autoincr spu_id" json:"spuId"`
	AttrName  string `xorm:"varchar(30) attr_name" json:"attrName"`
	AttrValue string `xorm:"text attr_value" json:"attrValue"`
}

// 设定表名
func (t *Invoice) TableName() string {
	return "invoice"
}
