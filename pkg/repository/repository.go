package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
)

type CarList interface {
	Create(list Drom.Car) (int, error)
	GetAll() ([]Drom.Car, error)
	GetById(listId int) (Drom.Car, error)
	Delete(listId int) error
	Update(lisId int, input Drom.UpdateListInput) error
}
type CarItem interface {
}
type Repository struct {
	CarItem
	CarList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CarList: NewCarsListPostgres(db),
	}
}
