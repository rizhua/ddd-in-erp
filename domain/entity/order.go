package entity

import (
	"time"

	"rizhua.com/domain/repository"
)

// 地址
type addressVO struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// 采购、销售商品
type itemVO struct {
	SpuName string
	SpuMeta string
	Sku     string
	Price   time.Time
	Amount  int8
	Deduct  string
}

var OrderRepo repository.Order

// 订单
type Order struct {
	ID             int64
	OrderNo        string
	TradeNo        string
	SupplierID     string
	SupplierName   string
	Amount         string
	Money          string
	PaymentChannel string
	PaymentAmount  string
	PaymentTime    string
	WaybillNO      string
	ConsigneeName  string
	ConsigneeTel   string
	ConsigneeAddr  *addressVO
	DeliveryMode   string
	ShipmentType   string
	ShipmentFee    string
	ShipmentRoute  string
	ConfirmTime    string
	ReceiveMode    string
	Status         string
	Currency       string
	PurchaseID     string
	PurchaseName   string
	Remark         string
	OrgID          int32
	UpdateAt       time.Time
	CreateAt       time.Time
	Item           []*itemVO
}

// 退换/售后
type RefundEntity struct {
	ID         int32
	RefundNo   string
	Data       []interface{}
	Reason     string
	Detail     string
	Amount     int32
	VerifyTime time.Time
	VerifyText string
	Status     int8
	OrderNo    string
	UpdateAt   time.Time
	CreateAt   time.Time
}
