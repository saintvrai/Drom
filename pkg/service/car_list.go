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
