package repository

type Order struct {
	ID       string  `json:"id"`
	Products string  `json:"name"`
	Total    float64 `json:"total"`
}
