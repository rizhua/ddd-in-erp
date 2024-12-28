package application

import (
	"context"
	"encoding/json"
	"errors"

	"rizhua.com/application/command"
	"rizhua.com/domain"
)

func NewOrderService(
	orderDomain domain.OrderService,
	userDomain domain.UserService,
) OrderService {
	return OrderService{
		orderDomain: orderDomain,
		userDomain:  userDomain,
	}
}

type OrderService struct {
	Context     context.Context
	orderDomain domain.OrderService
	userDomain  domain.UserService
}

func (t *OrderService) Unified(args []byte) (any, error) {
	// 1. 参数校验
	cmd := command.UnifiedOrder{}
	if err := json.Unmarshal(args, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// 2. 根据订单类型生成预订单
	switch cmd.Type {
	case command.OrderTypePack:

		return nil, nil
	case command.OrderTypeProduct:
		return nil, nil
	default:
		return nil, errors.New("未知的订单类型")
	}

	// 3. 接收支付消息

	// 4. 接收发货消息
}
