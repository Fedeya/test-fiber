package models

// Product is a simple type for products
type Product struct {
	Model
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint    `json:"quantity"`
}
