package service

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/repository"
)

type Car interface {
	Create(car Drom.Car) (int, error)
	GetAll() ([]Drom.Car, error)
	GetById(carId int) (Drom.Car, error)
	Delete(carId int) error
	Update(carId int, input Drom.UpdateListInput) error
}

type Service struct {
	Car
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Car: NewCarsService(repos.Car),
	}
}
