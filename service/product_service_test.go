package service

import (
	"ecommerce/model/entity"
	"ecommerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Register Repository
var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}

// Register Services
var productSvc = NewProductService(productRepository)

func TestSaveProduct(t *testing.T) {
	product := entity.Products{
		ID:      "123",
		SKU:     "S123",
		Name:    "Test Product",
		Stock:   1200,
		Price:   99000,
		Uom:     "Pcs",
		Images:  "",
		BrandID: "B81",
	}

	productRepository.Mock.On("SaveProducts").Return(nil)
	err := productSvc.SaveProduct(product)
	assert.Nil(t, err, "Should be nil")
}

func TestGetProduct(t *testing.T) {
	products := []entity.ProductResponse{
		{
			ID:        "123",
			SKU:       "S123",
			Name:      "Test Product",
			Stock:     1200,
			Price:     99000,
			Uom:       "Pcs",
			Images:    "",
			BrandID:   "B81",
			BrandName: "Test",
		},
		{
			ID:        "890",
			SKU:       "S12003",
			Name:      "Test Product 2",
			Stock:     1200,
			Price:     1200,
			Uom:       "Pcs",
			Images:    "",
			BrandID:   "B81",
			BrandName: "Test2",
		},
	}

	productRepository.Mock.On("GetProducts").Return(products)
	res, err := productSvc.GetProducts()
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, int(2), len(res), "Should be 2 items")
	assert.Equal(t, "Test Product 2", res[1].Name, "Name must be Test Product 2")
}

func TestGetProductByBrand(t *testing.T) {
	products := []entity.ProductResponse{
		{
			ID:        "123",
			SKU:       "S123",
			Name:      "Test Product",
			Stock:     1200,
			Price:     99000,
			Uom:       "Pcs",
			Images:    "",
			BrandID:   "B81",
			BrandName: "Test",
		},
		{
			ID:        "890",
			SKU:       "S12003",
			Name:      "Test Product 2",
			Stock:     1200,
			Price:     1200,
			Uom:       "Pcs",
			Images:    "",
			BrandID:   "B81",
			BrandName: "Test2",
		},
	}

	productRepository.Mock.On("GetProductByBrand").Return(products)
	res, err := productSvc.GetProductByBrand("123")
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, int(2), len(res), "Should be 2 items")
	assert.Equal(t, "Test Product 2", res[1].Name, "Name must be Test Product 2")
}
