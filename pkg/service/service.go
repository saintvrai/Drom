package service

import (
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/internal/user"
	"github.com/saintvrai/Drom/pkg/repository"
)

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}
type Car interface {
	Create(car car.Car) (int, error)
	GetAll() ([]car.Car, error)
	GetById(carId int) (car.Car, error)
	Delete(carId int) error
	Update(carId int, input car.UpdateListInput) error
}

type Service struct {
	Car
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewCarsService(repos.Car),
		NewAuthService(repos.Authorization),
	}
}
