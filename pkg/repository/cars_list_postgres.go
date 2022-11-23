package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom"
	"github.com/sirupsen/logrus"
	"strings"
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

	createListQuery := fmt.Sprintf("INSERT INTO %s (NAME,CARBRAND) VALUES ($1,$2) RETURNING id", carsListTable)
	row := tx.QueryRow(createListQuery, list.Name, list.CarBrand)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *CarsListPostgres) GetAll() ([]Drom.Car, error) {
	var lists []Drom.Car
	query := fmt.Sprintf("SELECT * FROM %s", carsListTable)
	err := r.db.Select(&lists, query)
	return lists, err
}
func (r *CarsListPostgres) GetById(listId int) (Drom.Car, error) {
	var list Drom.Car
	query := fmt.Sprintf("SELECT * FROM %s tl WHERE tl.id = $1", carsListTable)
	err := r.db.Get(&list, query, listId)
	return list, err
}
func (r *CarsListPostgres) Delete(listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1", carsListTable)
	_, err := r.db.Exec(query, listId)
	return err
}
func (r *CarsListPostgres) Update(listId int, input Drom.UpdateListInput) error {
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
		carsListTable, setQuery, carsListTable, argId)
	args = append(args, listId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
