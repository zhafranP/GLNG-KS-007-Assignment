package order_repository

import "ddd/entity"

type Repository interface {
	CreateOrder(orderPayload entity.Order, itemPayload []entity.Item) error
	ReadOrders() ([]OrderItemMapped, error)
	UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*OrderItemMapped, error)
	DeleteOrder(orderId int) error
}
