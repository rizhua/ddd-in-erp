package domain

import (
	"errors"

	"rizhua.com/application/command"
	"rizhua.com/domain/entity"
	"rizhua.com/domain/repository"
)

func NewOrderService(orderRepo repository.Order) OrderService {
	entity.OrderRepo = orderRepo
	return OrderService{
		orderRepo: orderRepo,
	}
}

type OrderService struct {
	orderRepo repository.Order
}

// 统一下单
func (t *OrderService) Unified(cmd command.UnifiedOrder) (orderID int64, err error) {
	// 根据订单类型生成订单
	switch cmd.Type {
	case command.OrderTypePack:
		// nodeDO := entity.NodeEntity{}
		// orderID, err = nodeDO.CreateOrder(cmd.BizID, cmd.Qty)
		return
	case command.OrderTypeProduct:
		return
	default:
		err = errors.New("未知的订单类型")
		return
	}
}
