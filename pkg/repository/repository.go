package repository

type CarList interface {
}
type CarItem interface {
}
type Repository struct {
	CarItem
	CarList
}

func NewRepository() *Repository {
	return &Repository{}
}
