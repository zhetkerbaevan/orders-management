package models

import (
	"context"

	"github.com/zhetkerbaevan/orders-management/internal/genproto"
)

type OrderService interface {
	CreateOrder(context.Context, *genproto.Order) error
	GetOrders(context.Context) []*genproto.Order
}
