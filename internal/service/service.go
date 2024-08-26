package service

import (
	"context"

	"github.com/zhetkerbaevan/orders-management/internal/genproto"
)

var orders = make([]*genproto.Order, 0)

type OrderService struct {
	//Store dependency injection
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *genproto.Order) error {
	orders = append(orders, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*genproto.Order {
	return orders
}
