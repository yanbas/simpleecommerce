package repository

import (
	"ecommerce/model/entity"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	Mock mock.Mock
}

func (repo *OrderRepositoryMock) CheckStock(id string) (int, error) {
	return 100, nil
}

func (repo *OrderRepositoryMock) CheckProductPrice(id string) (float64, error) {
	return 100, nil
}

func (repo *OrderRepositoryMock) OrderProcess(entity.Order) error {
	return nil
}

func (repo *OrderRepositoryMock) OrderItemProcess(entity.OrderItem) error {
	return nil
}

func (repo *OrderRepositoryMock) DecreaseQty(productId string, amount int) error {
	return nil
}

func (repo *OrderRepositoryMock) GetOrder(id string) (entity.Order, error) {
	arguments := repo.Mock.Called()
	result := arguments.Get(0).(entity.Order)
	return result, nil
}

func (repo *OrderRepositoryMock) GetOrderItem(id string) ([]entity.OrderItemResponse, error) {
	arguments := repo.Mock.Called()
	result := arguments.Get(0).([]entity.OrderItemResponse)
	return result, nil
}
