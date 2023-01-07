package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom/internal/car"
	//"github.com/saintvrai/Drom/pkg/logging"
	log "github.com/sirupsen/logrus"
	"strings"
)

type CarsPostgres struct {
	db *sqlx.DB
}

func NewCarsPostgres(db *sqlx.DB) *CarsPostgres {
	return &CarsPostgres{db: db}
}
func (r *CarsPostgres) Create(car car.Car) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	var id string
	createListQuery := fmt.Sprintf("INSERT INTO %s (name,carbrand,free,clientid) VALUES ($1,$2,$3,$4) RETURNING id", carsTable)
	row := tx.QueryRow(createListQuery, car.Name, car.CarBrand, car.Free, car.Client.ID)
	if err := row.Scan(&id); err != nil {
		if err := tx.Rollback(); err != nil {

			log.Error(err)
			return "", err
		}
		return "", err
	}
	return id, tx.Commit()
}
func (r *CarsPostgres) GetAll() ([]car.Car, error) {
	var cars []car.Car
	query := fmt.Sprintf("SELECT * FROM %s", carsTable)
	err := r.db.Select(&cars, query)
	return cars, err
}
func (r *CarsPostgres) GetById(carId int) (car.Car, error) {
	var carObj car.Car
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", carsTable)
	err := r.db.Get(&carObj, query, carId)
	return carObj, err
}
func (r *CarsPostgres) Delete(carId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1", carsTable)
	_, err := r.db.Exec(query, carId)
	return err
}
func (r *CarsPostgres) Update(carId int, input car.UpdateListInput) error {
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
	args = append(args, carId)

	log.Debugf("updateQuery: %s", query)
	log.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}

func (r *CarsPostgres) GetAllCarsAndClients() (list []car.CarAndClientName, err error) {
	var carAndClientName car.CarAndClientName
	var carAndClientNames []car.CarAndClientName
	createListQuery := fmt.Sprintf("SELECT crs.name, cls.name FROM %s crs INNER JOIN %s cls on crs.clientid = cls.id", carsTable, clientTable)
	rows, err := r.db.Queryx(createListQuery)
	if err != nil {

	}
	for rows.Next() {
		if err := rows.Scan(&carAndClientName.CarName, &carAndClientName.ClientName); err != nil {
			log.Errorf(err.Error())
			return nil, err
		}
		carAndClientNames = append(carAndClientNames, carAndClientName)
	}
	return carAndClientNames, nil
}
