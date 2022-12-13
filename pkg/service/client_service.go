package service

import (
	"github.com/saintvrai/Drom/internal/client"
	"github.com/saintvrai/Drom/pkg/repository"
)

type ClientsService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientsService {
	return &ClientsService{repo: repo}
}
func (s *ClientsService) Create(client client.Client) (string, error) {
	return s.repo.Create(client)
}
func (s *ClientsService) GetAll() ([]client.Client, error) {
	return s.repo.GetAll()
}
func (s *ClientsService) GetById(clientId string) (client.Client, error) {
	return s.repo.GetById(clientId)
}
func (s *ClientsService) Delete(clientId string) error {
	return s.repo.Delete(clientId)
}
func (s *ClientsService) Update(clientId string, input client.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(clientId, input)

}
