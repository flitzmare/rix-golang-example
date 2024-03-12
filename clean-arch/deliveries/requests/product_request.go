package requests

type CreateProduct struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID uint    `json:"user_id"`
}
