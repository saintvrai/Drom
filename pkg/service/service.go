package service

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/repository"
)

type CarList interface {
	Create(car Drom.Car) (int, error)
}
type CarItem interface {
}
type Service struct {
	CarItem
	CarList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{CarList: NewCarsListService(repos.CarList)}
}
