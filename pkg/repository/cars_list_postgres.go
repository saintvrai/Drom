package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
)

type CarsListPostgres struct {
	db *sqlx.DB
}

func NewCarsListPostgres(db *sqlx.DB) *CarsListPostgres {
	return &CarsListPostgres{db: db}
}
func (r *CarsListPostgres) Create(list Drom.Car) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (carbrandname,name) VALUES ($1,$2) RETURNING id", carsListTable)
	row := tx.QueryRow(createListQuery, list.CarBrand, list.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
