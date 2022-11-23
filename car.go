package Drom

import "errors"

type Car struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	CarBrand string `json:"carbrand" db:"carbrand"`
}
type UpdateListInput struct {
	Name     *string `json:"name"`
	CarBrand *string `json:"carbrand"`
}

func (i UpdateListInput) Validate() error {
	if i.Name == nil && i.CarBrand == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
