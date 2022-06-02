package entity

type OrderItem struct {
	ID        string  `json:"id,omitempty"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	Promo     float64 `json:"promo"`
	Total     float64 `json:"total"`
}
