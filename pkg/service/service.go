package service

import "github.com/saintvrai/Drom/pkg/repository"

type CarList interface {
}
type CarItem interface {
}
type Service struct {
	CarItem
	CarList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
