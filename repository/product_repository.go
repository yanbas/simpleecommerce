package repository

import "ecommerce/model/entity"

type ProductRepository interface {
	SaveProducts(entity.Products) error
	GetProducts() ([]entity.ProductResponse, error)
	GetProductByBrand(id string) ([]entity.ProductResponse, error)
}
