package registry

import (
	"database/sql"
	"ecommerce/repository"
	"ecommerce/service"
)

func RegisterProductService(db *sql.DB) service.ProductService {
	return service.NewProductService(
		repository.NewProductRepository(db),
	)
}

func RegisterOrderService(db *sql.DB) service.OrderService {
	return service.NewOrderService(
		repository.NewOrderRepository(db),
	)
}
