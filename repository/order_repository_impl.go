package repository

import (
	"database/sql"
	"ecommerce/model/entity"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepository(conn *sql.DB) OrderRepository {
	return &orderRepo{db: conn}
}

func (p *orderRepo) OrderProcess(order entity.Order) error {
	_, err := p.db.Exec(`INSERT INTO "order"(id, order_date, total, promo, subtotal) VALUES($1, $2, $3, $4, $5)`,
		order.ID, order.OrderDate, order.Total, order.Promo, order.SubTotal)

	if err != nil {
		return err
	}

	return nil
}

func (p *orderRepo) OrderItemProcess(order entity.OrderItem) error {
	_, err := p.db.Exec(`INSERT INTO order_item(id, order_id, product_id, price, qty, promo, total) 
						VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		order.ID, order.OrderID, order.ProductID, order.Price, order.Qty, order.Promo, order.Total)

	if err != nil {
		return err
	}

	return nil
}

func (p *orderRepo) CheckStock(id string) (int, error) {
	var stock int
	err := p.db.QueryRow(`SELECT stock FROM products WHERE id = $1`, id).Scan(&stock)
	if err != nil {
		return stock, err
	}
	return stock, nil
}

func (p *orderRepo) CheckProductPrice(id string) (float64, error) {
	var price float64
	err := p.db.QueryRow(`SELECT price FROM products WHERE id = $1`, id).Scan(&price)
	if err != nil {
		return price, err
	}
	return price, nil
}

func (p *orderRepo) DecreaseQty(id string, amount int) error {
	_, err := p.db.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2`, amount, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *orderRepo) GetOrder(id string) (entity.Order, error) {
	var o entity.Order
	err := p.db.QueryRow(`SELECT id, order_date, total, promo, subtotal FROM "order" WHERE id = $1`, id).Scan(&o.ID, &o.OrderDate, &o.Total, &o.Promo, &o.SubTotal)
	if err != nil {
		return o, err
	}
	return o, nil

}

func (p *orderRepo) GetOrderItem(id string) ([]entity.OrderItemResponse, error) {
	rows, err := p.db.Query(`SELECT id, product_id, price, qty, promo, total FROM order_item 
							WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderItem := []entity.OrderItemResponse{}

	for rows.Next() {
		var p entity.OrderItemResponse
		if err := rows.Scan(&p.ID, &p.ProductID, &p.Price, float64(p.Qty), &p.Promo, &p.Total); err != nil {
			return nil, err
		}

		orderItem = append(orderItem, p)
	}

	return orderItem, nil
}
