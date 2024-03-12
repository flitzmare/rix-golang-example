package responses

type GetProduct struct {
	ID    uint    `json:"id" example:"10"`
	Name  string  `json:"name" example:"Riksa"`
	Price float64 `json:"price" example:"20000"`
}
