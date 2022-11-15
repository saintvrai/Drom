package repository

import (
	"github.com/jmoiron/sqlx"
)

type CarList interface {
}
type CarItem interface {
}
type Repository struct {
	CarItem
	CarList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
