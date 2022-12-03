package service

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/repository"
)

type CarsService struct {
	repo repository.Car
}

func NewCarsService(repo repository.Car) *CarsService {
	return &CarsService{repo: repo}
}
func (s *CarsService) Create(car Drom.Car) (int, error) {
	return s.repo.Create(car)
}
func (s *CarsService) GetAll() ([]Drom.Car, error) {
	return s.repo.GetAll()
}
func (s *CarsService) GetById(carId int) (Drom.Car, error) {
	return s.repo.GetById(carId)
}
func (s *CarsService) Delete(carId int) error {
	return s.repo.Delete(carId)
}
func (s *CarsService) Update(carId int, input Drom.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(carId, input)

}
