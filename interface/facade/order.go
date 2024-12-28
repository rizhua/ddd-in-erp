package facade

import (
	"rizhua.com/application"
)

type OrderFacade interface {
}

func NewOrderFacade(productApp application.ProductService) OrderFacade {
	return &orderFacade{
		productApp: productApp,
	}
}

type orderFacade struct {
	productApp application.ProductService // 商品订单
}

// 统一下单
// @param typ 1商品订单 2套餐订单
func (t *orderFacade) Unified(typ int, args []byte) {
	switch typ {
	case 1:
		// t.productApp.CreateOrder(context.Background(), cmd)
	case 2:
		// t.nodeApp.CreateOrder(args)
	}
}
