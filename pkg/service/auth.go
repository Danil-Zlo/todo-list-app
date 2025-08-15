package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Danil-Zlo/todo-list-app"
	"github.com/Danil-Zlo/todo-list-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "74f3s3d63awho4rgta39wejh63bct4221fwcr4n09owt7n"
	tokenTTL   = 12 * time.Hour
	signingKey = "jo2uh6572pktrfs254gaeerr3"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// первый аргумент - подпись
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			// действие токена 12 часов
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	// create hash
	hash := sha1.New()
	hash.Write([]byte(password))

	// return hash with salt
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
