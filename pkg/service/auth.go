package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Danil-Zlo/todo-list-app"
	"github.com/Danil-Zlo/todo-list-app/pkg/repository"
)

const salt = "74f3s3d63awho4rgta39wejh63bct4221fwcr4n09owt7n"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	// create hash
	hash := sha1.New()
	hash.Write([]byte(password))

	// return hash with salt
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
