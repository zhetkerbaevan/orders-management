package service

import (
	"context"

	"github.com/zhetkerbaevan/orders-management/internal/genproto"
	"github.com/zhetkerbaevan/orders-management/internal/models"
)

type OrderService struct {
	orderStore models.OrderStore //Store dependency injection
}

func NewOrderService(orderStore models.OrderStore) *OrderService {
	return &OrderService{orderStore: orderStore}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *genproto.Order) error {
	err := s.orderStore.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*genproto.Order {
	os, err := s.orderStore.GetOrders()
	if err != nil {
		return nil
	}

	return os
}

func (s *OrderService) DeleteOrder(ctx context.Context, id int32) error {
	err := s.orderStore.DeleteOrder(id)
	if err != nil {
		return err
	}

	return nil
}
