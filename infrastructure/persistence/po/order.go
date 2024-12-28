package po

import "time"

// 订单
type Order struct {
	ID             int64     `xorm:"id pk autoincr"`
	OrderNo        string    `xorm:"order_no varchar(90) notnull"`
	TradeNo        string    `xorm:"trade_no varchar(90) notnull"`
	SupplierID     string    `xorm:"supplier_id varchar(90) notnull"`
	SupplierName   string    `xorm:"supplier_name varchar(90) notnull"`
	Amount         string    `xorm:"amount varchar(90) notnull"`
	Money          string    `xorm:"money varchar(90) notnull"`
	PaymentChannel string    `xorm:"payment_channel varchar(90) notnull"`
	PaymentAmount  string    `xorm:"payment_amount varchar(90) notnull"`
	PaymentTime    string    `xorm:"payment_time varchar(90) notnull"`
	WaybillNO      string    `xorm:"waybill_no varchar(90) notnull"`
	ConsigneeName  string    `xorm:"consignee_name varchar(90) notnull"`
	ConsigneeTel   string    `xorm:"consignee_tel varchar(90) notnull"`
	ConsigneeAddr  string    `xorm:"consignee_addr varchar(90) notnull"`
	DeliveryMode   string    `xorm:"delivery_mode varchar(90) notnull"`
	ShipmentType   string    `xorm:"shipment_type varchar(90) notnull"`
	ShipmentFee    string    `xorm:"shipment_fee varchar(90) notnull"`
	ShipmentRoute  string    `xorm:"shipment_route varchar(90) notnull"`
	ConfirmTime    string    `xorm:"confirm_time varchar(90) notnull"`
	ReceiveMode    string    `xorm:"receive_mode varchar(90) notnull"`
	Status         string    `xorm:"status varchar(90) notnull"`
	Currency       string    `xorm:"currency varchar(90) notnull"`
	PurchaseID     string    `xorm:"purchase_id varchar(90) notnull"`
	PurchaseName   string    `xorm:"purchase_name varchar(90) notnull"`
	Remark         string    `xorm:"remark varchar(90) notnull"`
	OrgID          int32     `xorm:"org_id int"`
	UpdateAt       time.Time `xorm:"update_at datatime"`
	CreateAt       time.Time `xorm:"create_at datetime"`
}

// 设定表名
func (t *Order) TableName() string {
	return "order"
}
