package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom/internal/car"
)

type Car interface {
	Create(car car.Car) (int, error)
	GetAll() ([]car.Car, error)
	GetById(carId int) (car.Car, error)
	Delete(carId int) error
	Update(carId int, input car.UpdateListInput) error
}
type Repository struct {
	Car
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Car: NewCarsPostgres(db),
	}
}
