package models

type Brand struct {
	ID    int    `valid:"-" json:"id" db:"id"`
	Name  string `valid:"-" json:"name" db:"brand_name"`
	Year  int    `valid:"-" json:"year" db:"founding_year"`
	Logo  int    `valid:"-" json:"logo" db:"logo_id"`
	Owner string `valid:"-" json:"owner" db:"brand_owner"`
}
