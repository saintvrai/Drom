package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/saintvrai/Drom/internal/client"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ClientsPostgres struct {
	db *sqlx.DB
}

func NewClientsPostgres(db *sqlx.DB) *ClientsPostgres {
	return &ClientsPostgres{db: db}
}
func (r *ClientsPostgres) Create(client client.Client) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	var id string
	createListQuery := fmt.Sprintf("INSERT INTO %s (NAME,MONEY) VALUES ($1,$2) RETURNING id", clientTable)
	row := tx.QueryRow(createListQuery, client.Name, client.Money)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return "", err
		}
		return "", err
	}
	return id, tx.Commit()
}
func (r *ClientsPostgres) GetAll() ([]client.Client, error) {
	var clients []client.Client
	query := fmt.Sprintf("SELECT * FROM %s", clientTable)
	err := r.db.Select(&clients, query)
	return clients, err
}
func (r *ClientsPostgres) GetById(clientID string) (client.Client, error) {
	var clientObj client.Client
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", clientTable)
	err := r.db.Get(&clientObj, query, clientID)
	return clientObj, err
}
func (r *ClientsPostgres) Delete(clientID string) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1", clientTable)
	_, err := r.db.Exec(query, clientID)
	return err
}
func (r *ClientsPostgres) Update(clientID string, input client.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Money != nil {
		setValues = append(setValues, fmt.Sprintf("carbrand=$%d", argId))
		args = append(args, *input.Money)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s WHERE tl.id=$%d",
		clientTable, setQuery, clientTable, argId)
	args = append(args, clientID)

	log.Debugf("updateQuery: %s", query)
	log.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
