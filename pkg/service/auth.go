package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/saintvrai/Drom/internal/user"
	"github.com/saintvrai/Drom/pkg/repository"
)

const salt = "2001vsemprivet2022"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user user.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
