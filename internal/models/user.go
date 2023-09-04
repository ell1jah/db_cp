package models

type User struct {
	ID       int    `valid:"-" json:"id" db:"user_id"`
	Login    string `valid:"minstringlength(5)" json:"login" db:"user_login"`
	Password string `valid:"minstringlength(5)" json:"password" db:"user_password"`
	Name     string `valid:"minstringlength(2)" json:"name" db:"user_name"`
	Sex      string `valid:"in(male|female)" json:"sex" db:"user_sex"`
	Role     string `valid:"in(admin|guest|user)" json:"role" db:"user_role"`
}
