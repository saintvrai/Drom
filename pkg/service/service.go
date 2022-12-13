package service

import (
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/internal/client"
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
type Client interface {
	Create(client client.Client) (string, error)
	GetAll() ([]client.Client, error)
	GetById(clientId string) (client.Client, error)
	Delete(clientId string) error
	Update(clientId string, input client.UpdateListInput) error
}
type Service struct {
	Car
	Client
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewCarsService(repos.Car),
		NewClientService(repos.Client),
		NewAuthService(repos.Authorization),
	}
}
