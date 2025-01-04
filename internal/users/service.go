package users

import (
	"github.com/vferreirati/go_bookstore/internal/auth"
	"github.com/vferreirati/go_bookstore/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(name string, email string, password string) (models.User, error)
	Login(email string, password string) (models.Login, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateUser(name string, email string, password string) (models.User, error) {
	user, err := s.repo.CreateUser(name, email, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *service) Login(email string, password string) (models.Login, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return models.Login{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.Login{}, err
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return models.Login{}, err
	}
	return models.Login{Token: token, UserID: user.ID}, nil
}
