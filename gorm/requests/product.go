package requests

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PaginationRequest struct {
	Size int `query:"size"`
	Page int `query:"page"`
}

type PaginationResponse struct {
	TotalData  int64       `json:"total_data"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
