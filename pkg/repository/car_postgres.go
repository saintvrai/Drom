package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
	"github.com/sirupsen/logrus"
	"strings"
)

type CarsPostgres struct {
	db *sqlx.DB
}

func NewCarsPostgres(db *sqlx.DB) *CarsPostgres {
	return &CarsPostgres{db: db}
}
func (r *CarsPostgres) Create(car Drom.Car) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (NAME,CARBRAND) VALUES ($1,$2) RETURNING id", carsTable)
	row := tx.QueryRow(createListQuery, car.Name, car.CarBrand)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	return id, tx.Commit()
}
func (r *CarsPostgres) GetAll() ([]Drom.Car, error) {
	var lists []Drom.Car
	query := fmt.Sprintf("SELECT * FROM %s", carsTable)
	err := r.db.Select(&lists, query)
	return lists, err
}
func (r *CarsPostgres) GetById(listId int) (Drom.Car, error) {
	var car Drom.Car
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", carsTable)
	err := r.db.Get(&car, query, listId)
	return car, err
}
func (r *CarsPostgres) Delete(listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1", carsTable)
	_, err := r.db.Exec(query, listId)
	return err
}
func (r *CarsPostgres) Update(listId int, input Drom.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.CarBrand != nil {
		setValues = append(setValues, fmt.Sprintf("carbrand=$%d", argId))
		args = append(args, *input.CarBrand)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s WHERE tl.id=$%d",
		carsTable, setQuery, carsTable, argId)
	args = append(args, listId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
