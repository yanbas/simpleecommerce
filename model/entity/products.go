package entity

type Products struct {
	ID      string  `json:"id,omitempty"`
	SKU     string  `json:"sku"`
	Name    string  `json:"name"`
	Uom     string  `json:"uom"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Images  string  `json:"images"`
	BrandID string  `json:"brand_id"`
}

type ProductResponse struct {
	ID        string  `json:"id,omitempty"`
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Uom       string  `json:"uom"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Images    string  `json:"images"`
	BrandID   string  `json:"brand_id"`
	BrandName string  `json:"brand_name"`
}
