package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
)

type Car interface {
	Create(car Drom.Car) (int, error)
	GetAll() ([]Drom.Car, error)
	GetById(carId int) (Drom.Car, error)
	Delete(carId int) error
	Update(carId int, input Drom.UpdateListInput) error
}
type Repository struct {
	Car
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Car: NewCarsPostgres(db),
	}
}
