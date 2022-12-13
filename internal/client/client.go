package client

import "errors"

type Client struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Money string `json:"money" db:"money"`
}
type UpdateListInput struct {
	Name  *string `json:"name" db:"name"`
	Money *int    `json:"money" db:"money"`
}

func (i UpdateListInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
