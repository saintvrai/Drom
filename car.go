package Drom

type Car struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"id"`
	CarBrand string `json:"carBrand" db:"id"`
}
