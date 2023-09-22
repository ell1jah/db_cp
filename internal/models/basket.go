package models

type Basket struct {
	ID    int         `valid:"-" json:"id" db:"id"`
	Items []OrderItem `valid:"-" json:"items" db:"items"`
	Price int         `valid:"-" json:"price" db:"price"`
}

func NewBasket() *Basket {
	return &Basket{
		Items: []OrderItem{},
	}
}
