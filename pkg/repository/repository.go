package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/internal/client"
	"github.com/saintvrai/Drom/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) (string, error)
	GetUser(username, password string) (user.User, error)
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
type Repository struct {
	Authorization
	Client
	Car
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Car:           NewCarsPostgres(db),
		Client:        NewClientsPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
