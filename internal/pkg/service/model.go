package service

type Order struct {
	ID       string  `json:"id" db:"id"`
	Products string  `json:"name" db:"products"`
	Total    float64 `json:"total" db:"total"`
}
