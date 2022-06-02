package repository

import "ecommerce/model/entity"

type OrderRepository interface {
	CheckStock(id string) (int, error)
	CheckProductPrice(id string) (float64, error)
	OrderProcess(entity.Order) error
	OrderItemProcess(entity.OrderItem) error
	DecreaseQty(productId string, amount int) error
	GetOrder(id string) (entity.Order, error)
	GetOrderItem(id string) ([]entity.OrderItemResponse, error)
}
