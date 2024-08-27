package store

import (
	"database/sql"

	"github.com/zhetkerbaevan/orders-management/internal/genproto"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{db: db}
}

func (s *OrderStore) CreateOrder(order *genproto.Order) error {
	_, err := s.db.Exec("INSERT INTO orders (item, quantity) VALUES ($1, $2)", order.Item, order.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderStore) GetOrders() ([]*genproto.Order, error) {
	rows, err := s.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}

	orders := make([]*genproto.Order, 0)

	for rows.Next() {
		o, err := ScanIntoOrders(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (s *OrderStore) DeleteOrder(id int32) error {
	_, err := s.db.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func ScanIntoOrders(rows *sql.Rows) (*genproto.Order, error) {
	order := new(genproto.Order)

	err := rows.Scan(
		&order.Id,
		&order.Item,
		&order.Quantity)
	if err != nil {
		return nil, err
	}

	return order, nil
}
