package handler

import (
	"context"

	"github.com/utrack/clay/v3/transport"
	"github.com/zhetkerbaevan/orders-management/internal/genproto"
	"github.com/zhetkerbaevan/orders-management/internal/models"
)

type OrderImplementation struct {
	orderService models.OrderService
	genproto.UnimplementedOrderServiceServer
}

func NewOrderImplementation(orderService models.OrderService) *OrderImplementation {
	return &OrderImplementation{orderService: orderService}
}

func (h *OrderImplementation) GetDescription() transport.ServiceDesc {
	return genproto.NewOrderServiceServiceDesc(h)
}

//Implementation of RPC methods

func (h *OrderImplementation) CreateOrder(ctx context.Context, req *genproto.CreateOrderRequest) (*genproto.CreateOrderResponse, error) {
	order := &genproto.Order{ //Get payload
		Item:     req.Item,
		Quantity: req.Quantity,
	}

	err := h.orderService.CreateOrder(ctx, order) //Send payload to service
	if err != nil {
		return nil, err
	}

	res := &genproto.CreateOrderResponse{
		Status: "Created",
	}

	return res, nil
}

func (h *OrderImplementation) GetOrder(ctx context.Context, req *genproto.CetOrderRequest) (*genproto.GetOrderResponse, error) {
	os := h.orderService.GetOrders(ctx)
	res := &genproto.GetOrderResponse{
		Orders: os,
	}

	return res, nil
}
