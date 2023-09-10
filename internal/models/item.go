package models

type Item struct {
	ID          int    `valid:"-" json:"id" db:"id"`
	Category    string `valid:"in(ботинки|кроссовки|майка|футболка|куртка|штаны|шорты|ремень|шляпа)" json:"category" db:"category"`
	Size        string `valid:"in(XS|S|M|L|XL|XXL)" json:"size" db:"size"`
	Price       int    `valid:"-" json:"price" db:"price"`
	Sex         string `valid:"in(male|female)" json:"sex" db:"sex"`
	ImageID     int    `valid:"-" json:"image_id" db:"image_id"`
	BrandID     int    `valid:"-" json:"brand_id" db:"brand_id"`
	IsAvailable bool   `valid:"-" json:"is_available" db:"is_available"`
}

const (
	ItemsParamsAny = "any"
	ItemsOrderDesc = "desc"
	ItemsOrderAsc  = "asc"
)

type ItemsParams struct {
	Category string `valid:"in(ботинки|кроссовки|майка|футболка|куртка|штаны|шорты|ремень|шляпа|any)" json:"category"`
	Sex      string `valid:"in(male|female|any)" json:"sex"`
	Brand    int    `valid:"-" json:"brand"`
	Order    string `valid:"in(asc|desc|any)" json:"order"`
}
