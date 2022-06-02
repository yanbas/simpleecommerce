package service

import (
	"ecommerce/model/entity"
	"ecommerce/repository"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	Checkout(entity.OrderRequest) error
	GetOrderDetail(id string) (entity.OrderResponse, error)
}

type orderServiceImpl struct {
	Repository repository.OrderRepository
}

func NewOrderService(
	pr repository.OrderRepository,
) OrderService {
	return &orderServiceImpl{pr}
}

func (s *orderServiceImpl) Checkout(order entity.OrderRequest) error {
	orderdate := time.Now().Format("2006-01-02 15:04:05")
	for _, v := range order.Item {
		// check stock
		stock, err := s.Repository.CheckStock(v.ProductID)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if stock < v.Qty {
			return errors.New("Stock not enought")
		}
	}

	orderId := uuid.New().String()
	var totalPrice float64

	for _, i := range order.Item {
		productPrice, err := s.Repository.CheckProductPrice(i.ProductID)
		if err != nil {
			return err
		}

		orderDetail := entity.OrderItem{
			ID:        uuid.New().String(),
			OrderID:   orderId,
			ProductID: i.ProductID,
			Qty:       i.Qty,
			Promo:     0,
			Price:     productPrice,
			Total:     productPrice * float64(i.Qty),
		}

		err = s.Repository.OrderItemProcess(orderDetail)
		if err != nil {
			return err
		}

		// decrease stock
		err = s.Repository.DecreaseQty(i.ProductID, i.Qty)
		if err != nil {
			return err
		}

		totalPrice = +productPrice * float64(i.Qty)

	}

	orderData := entity.Order{
		ID:        orderId,
		OrderDate: orderdate,
		Promo:     0,
		SubTotal:  totalPrice,
		Total:     totalPrice,
	}

	err := s.Repository.OrderProcess(orderData)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderServiceImpl) GetOrderDetail(id string) (entity.OrderResponse, error) {
	order, err := s.Repository.GetOrder(id)
	if err != nil {
		return entity.OrderResponse{}, err
	}

	orderItem, err := s.Repository.GetOrderItem(id)
	if err != nil {
		return entity.OrderResponse{}, err
	}

	results := entity.OrderResponse{
		ID:        order.ID,
		OrderDate: order.OrderDate,
		Total:     order.Total,
		Promo:     order.Promo,
		SubTotal:  order.SubTotal,
		Items:     orderItem,
	}

	return results, nil
}
