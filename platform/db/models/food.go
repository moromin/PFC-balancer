package models

type Food struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbohydrate float64 `json:"carbohydrate"`
	Category     int64   `json:"category"`
}
