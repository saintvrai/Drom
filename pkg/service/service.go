package service

import (
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/pkg/repository"
)

type Car interface {
	Create(car car.Car) (int, error)
	GetAll() ([]car.Car, error)
	GetById(carId int) (car.Car, error)
	Delete(carId int) error
	Update(carId int, input car.UpdateListInput) error
}

type Service struct {
	Car
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Car: NewCarsService(repos.Car),
	}
}
