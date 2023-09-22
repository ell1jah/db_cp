package models

import "time"

type Order struct {
	ID     int         `valid:"-" json:"id" db:"id"`
	Date   time.Time   `valid:"-" json:"date" db:"commit_date"`
	UserID int         `valid:"-" json:"user" db:"user_id"`
	Items  []OrderItem `valid:"-" json:"items" db:"items"`
	Price  int         `valid:"-" json:"price" db:"price"`
	Status string      `valid:"-" json:"status" db:"current_status"`
}

func NewOrder() *Order {
	return &Order{
		Items: []OrderItem{},
	}
}
