package service

import (
	"ecommerce/model/entity"
	"ecommerce/repository"
)

type ProductService interface {
	SaveProduct(entity.Products) error
	GetProducts() ([]entity.ProductResponse, error)
	GetProductByBrand(brandID string) ([]entity.ProductResponse, error)
}

type productServiceImpl struct {
	Repository repository.ProductRepository
}

func NewProductService(
	pr repository.ProductRepository,
) ProductService {
	return &productServiceImpl{pr}
}

func (s *productServiceImpl) SaveProduct(product entity.Products) error {
	err := s.Repository.SaveProducts(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) GetProducts() ([]entity.ProductResponse, error) {
	result, err := s.Repository.GetProducts()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *productServiceImpl) GetProductByBrand(brandID string) ([]entity.ProductResponse, error) {
	result, err := s.Repository.GetProductByBrand(brandID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
