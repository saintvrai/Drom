package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
)

type CarList interface {
	Create(list Drom.Car) (int, error)
}
type CarItem interface {
}
type Repository struct {
	CarItem
	CarList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{CarList: NewCarsListPostgres(db)}
}
