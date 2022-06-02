package entity

type OrderRequest struct {
	UserID string             `json:"user_id"`
	Item   []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Qty       int    `json:"qty"`
}

type Order struct {
	ID        string  `json:"id"`
	OrderDate string  `json:"order_date"`
	Total     float64 `json:"total"`
	Promo     float64 `json:"promo"`
	SubTotal  float64 `json:"sub_total"`
}

type OrderItemResponse struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	Promo     float64 `json:"promo"`
	Total     float64 `json:"total"`
}

type OrderResponse struct {
	ID        string              `json:"id"`
	OrderDate string              `json:"order_date"`
	Total     float64             `json:"total"`
	Promo     float64             `json:"promo"`
	SubTotal  float64             `json:"sub_total"`
	Items     []OrderItemResponse `json:"items"`
}
