package Drom

type Car struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	CarBrand string `json:"carBrand"`
	CarColor string `json:"carColor"`
}
