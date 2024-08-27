package models

import (
	"context"

	"github.com/zhetkerbaevan/orders-management/internal/genproto"
)

type OrderService interface {
	CreateOrder(context.Context, *genproto.Order) error
	GetOrders(context.Context) []*genproto.Order
	DeleteOrder(context.Context, int32) error
}

type OrderStore interface {
	CreateOrder(*genproto.Order) error
	GetOrders() ([]*genproto.Order, error)
	DeleteOrder(int32) error
}
