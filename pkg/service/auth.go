package service

import (
	"crypto/sha1"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/saintvrai/Drom/internal/user"
	"github.com/saintvrai/Drom/pkg/repository"
	"time"
)

const salt = "2001vsemprivet2022"
const tokenTTL = 12 * time.Hour
const signingKey = "hwebfds7304sdnojb"

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_ID"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user user.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
