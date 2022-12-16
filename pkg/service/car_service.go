package service

import (
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/pkg/repository"
)

type CarsService struct {
	repo repository.Car
}

func NewCarsService(repo repository.Car) *CarsService {
	return &CarsService{repo: repo}
}
func (s *CarsService) Create(car car.Car) (string, error) {
	return s.repo.Create(car)
}
func (s *CarsService) GetAll() ([]car.Car, error) {
	return s.repo.GetAll()
}
func (s *CarsService) GetById(carId int) (car.Car, error) {
	return s.repo.GetById(carId)
}
func (s *CarsService) Delete(carId int) error {
	return s.repo.Delete(carId)
}
func (s *CarsService) Update(carId int, input car.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(carId, input)
}
func (s *CarsService) GetAllCarsAndClients() ([]car.Car, error) {
	return s.repo.GetAllCarsAndClients()
}
