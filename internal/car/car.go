package car

import (
	"errors"
	"github.com/saintvrai/Drom/internal/client"
)

type Car struct {
	ID       string        `json:"id" db:"id"`
	Name     string        `json:"name" db:"name"`
	CarBrand string        `json:"carbrand" db:"carbrand"`
	Free     bool          `json:"free" db:"free"`
	Client   client.Client `json:"client" db:"client"`
}
type UpdateListInput struct {
	Name     *string `json:"name"`
	CarBrand *string `json:"carbrand"`
}
type CarAndClientName struct {
	CarName    *string `json:"carname"`
	ClientName *string `json:"clientname"`
}

func (i UpdateListInput) Validate() error {
	if i.Name == nil && i.CarBrand == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
