package service

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/repository"
)

type CarList interface {
	Create(car Drom.Car) (int, error)
	GetAll() ([]Drom.Car, error)
	GetById(listId int) (Drom.Car, error)
	Delete(listId int) error
	Update(lisId int, input Drom.UpdateListInput) error
}
type CarItem interface {
}
type Service struct {
	CarItem
	CarList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CarList: NewCarsListService(repos.CarList),
	}
}
