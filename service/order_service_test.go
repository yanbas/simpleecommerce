package service

import (
	"ecommerce/model/entity"
	"ecommerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Register Repository
var orderRepository = &repository.OrderRepositoryMock{Mock: mock.Mock{}}

// Register Services
var orderSvc = NewOrderService(orderRepository)

func TestCheckout(t *testing.T) {
	order := entity.OrderRequest{
		UserID: "U100",
		Item: []entity.OrderItemRequest{
			{
				ProductID: "P111",
				Qty:       10,
			},
		},
	}

	productRepository.Mock.On("SaveProducts").Return(nil)
	err := orderSvc.Checkout(order)
	assert.Nil(t, err, "Should be nil")
}
