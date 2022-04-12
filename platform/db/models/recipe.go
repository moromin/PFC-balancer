package models

type FoodAmount struct {
	FoodId int64   `json:"food_amount"`
	Amount float64 `json:"amount"`
}

type Recipe struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	FoodAmounts []*FoodAmount `json:"food_amounts"`
	Procedures  []string      `json:"procedures"`
	UserId      int64         `json:"user_id"`
}
