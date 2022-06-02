package repository

import (
	"ecommerce/model/entity"
	"errors"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repo *ProductRepositoryMock) SaveProducts(product entity.Products) error {
	return nil
}

func (repo *ProductRepositoryMock) GetProducts() ([]entity.ProductResponse, error) {
	arguments := repo.Mock.Called()
	result := arguments.Get(0).([]entity.ProductResponse)
	return result, nil
}

func (repo *ProductRepositoryMock) GetProductByBrand(id string) ([]entity.ProductResponse, error) {
	if id == "123" {
		arguments := repo.Mock.Called()
		result := arguments.Get(0).([]entity.ProductResponse)
		return result, nil
	}
	return nil, errors.New("Brand ID not found")
}
