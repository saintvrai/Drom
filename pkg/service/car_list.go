package service

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/repository"
)

type CarsListService struct {
	repo repository.CarList
}

func NewCarsListService(repo repository.CarList) *CarsListService {
	return &CarsListService{repo: repo}
}
func (s *CarsListService) Create(list Drom.Car) (int, error) {
	return s.repo.Create(list)
}
func (s *CarsListService) GetAll() ([]Drom.Car, error) {
	return s.repo.GetAll()
}
func (s *CarsListService) GetById(listId int) (Drom.Car, error) {
	return s.repo.GetById(listId)
}
func (s *CarsListService) Delete(listId int) error {
	return s.repo.Delete(listId)
}
func (s *CarsListService) Update(lisId int, input Drom.UpdateListInput) error {
	return s.repo.Update(lisId, input)
}
