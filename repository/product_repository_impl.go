package repository

import (
	"database/sql"
	"ecommerce/model/entity"

	"github.com/google/uuid"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return &productRepo{db: conn}
}

func (p *productRepo) SaveProducts(product entity.Products) error {
	id := uuid.New().String()
	_, err := p.db.Exec(`INSERT INTO products(id, brand_id, name, uom, price, stock, images, sku) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		id, product.BrandID, product.Name, product.Uom, product.Price, product.Stock, product.Images, product.SKU)

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepo) GetProducts() ([]entity.ProductResponse, error) {
	rows, err := p.db.Query(`SELECT a.id, a.brand_id, a.name, a.uom, a.price, a.stock, a.images, a.sku, b."name" as brand_name
							FROM products as a
							join brand b on a.brand_id = b.id `)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []entity.ProductResponse{}

	for rows.Next() {
		var p entity.ProductResponse
		if err := rows.Scan(&p.ID, &p.BrandID, &p.Name, &p.Uom, &p.Price, &p.Stock, &p.Images, &p.SKU, &p.BrandName); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (p *productRepo) GetProductByBrand(id string) ([]entity.ProductResponse, error) {
	rows, err := p.db.Query(`SELECT a.id, a.brand_id, a.name, a.uom, a.price, a.stock, a.images, a.sku, b."name" as brand_name
							FROM products as a
							join brand b on a.brand_id = b.id 
							WHERE b.id = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []entity.ProductResponse{}

	for rows.Next() {
		var p entity.ProductResponse
		if err := rows.Scan(&p.ID, &p.BrandID, &p.Name, &p.Uom, &p.Price, &p.Stock, &p.Images, &p.SKU, &p.BrandName); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
