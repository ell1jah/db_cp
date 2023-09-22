package models

type OrderItem struct {
	Item
	Amount int `valid:"-" json:"amount" db:"amount"`
}
